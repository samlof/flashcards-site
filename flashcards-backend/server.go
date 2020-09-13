package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"flashcards-backend/auth"
	"flashcards-backend/ent"
	"flashcards-backend/ent/migrate"
	"flashcards-backend/graph"
	"flashcards-backend/graph/generated"
	"flashcards-backend/tts"
)

const defaultPort = "8080"

func main() {
	isProduction := applyDotEnv() == "production"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("No DATABASE_URL set")
	}

	// Make db client
	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed connecting to postgres: %v", err)
	}
	defer client.Close()

	// Run migrations
	ctx := context.Background()
	err = client.Schema.Create(ctx,
		//err = client.Schema.WriteTo(ctx, os.Stdout,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true))
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	srv := configureGqlServer(client)

	firebaseApp, err := setupFirebase()
	if err != nil {
		log.Fatalf("Error initializing firebase: %v", err)
	}
	firebaseAuth, err := firebaseApp.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing firebase auth: %v", err)
	}

	router := chi.NewRouter()

	router.With().HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			log.Printf("Error with healthcheck: %v", err)
		}
	})

	r := router.Group(nil)
	// A good base middleware stack
	//r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"authorization", "content-type", "accept"},
	}).Handler)

	authHandler := r.With(auth.Middleware(firebaseAuth, client))
	authHandler.Handle("/query", srv)
	r.Get("/tts/{text}-{lang}.mp3", ttsHandler())

	// Enable introspection and playground for development
	if !isProduction {
		srv.Use(extension.Introspection{})
		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground\n", port)
	}

	httpServer := http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Second * 5,
		IdleTimeout:       time.Minute * 1,
	}
	// Start server in goroutine, this one will wait for interrupt signal
	go func() {
		log.Printf("Listening at http://localhost:%s/query", port)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown code
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("Shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*9)
	defer cancel()

	err = httpServer.Shutdown(shutdownCtx)
	if err != nil {
		log.Printf("Graceful shutdown error: %v\n", err)
	} else {
		fmt.Println("Shutdown succesfully")
	}
}

func configureGqlServer(client *ent.Client) *handler.Server {
	rand.Seed(time.Now().UnixNano())
	graphResolver := &graph.Resolver{
		DB: client,
	}
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: graphResolver}))

	// Copied from handler.NewDefaultServer
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				// return r.Host == "example.org"
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			log.Printf("Got websocket init with payload: %v", initPayload)
			return nil, fmt.Errorf("Websocket not supported")
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	return srv
}

func ttsHandler() http.HandlerFunc {
	ttsService, err := tts.New(context.Background())
	if err != nil {
		log.Fatalf("Unable to make tts service: %v", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// Get params
		text := chi.URLParam(r, "text")
		lang := chi.URLParam(r, "lang")
		if text == "" || lang == "" {
			_, err = w.Write([]byte("params missing"))
			if err != nil {
				log.Printf("error writing params missing error: %v", err)
			}
		}
		text, err = url.QueryUnescape(text)
		if err != nil {
			log.Printf("error unescaping text: %v", err)
			_, err = w.Write([]byte("Invalid text param"))
			if err != nil {
				log.Printf("error writing invalid text param error: %v", err)
			}
			return
		}
		if lang == "en" {
			lang = "en-US"
		} else if lang == "fi" {
			lang = "fi-FI"
		} else {
			_, err = w.Write([]byte("Invalid language"))
			if err != nil {
				log.Printf("error writing invalid lang error: %v", err)
			}
			return
		}

		// Get audio bytes
		audio, err := ttsService.Tts(r.Context(), text, lang)
		if err != nil {
			log.Printf("error getting tts: %v", err)
			_, err = w.Write([]byte("error getting audio"))
			if err != nil {
				log.Printf("error writing error: %v", err)
			}
			return
		}

		w.Header().Add("content-type", "audio/mpeg3;audio/mpeg")
		// Cache the audio
		w.Header().Add("Cache-Control", "public, max-age=31536000")
		// Return bytes
		_, err = w.Write(audio)
		if err != nil {
			log.Printf("error writing tts: %v", err)
		}
	}
}

func applyDotEnv() string {
	env := os.Getenv("FC_ENV")
	if env == "" {
		env = "development"
	}
	_ = godotenv.Load(".env." + env + ".local")
	if env != "test" {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load()
	return env
}

func setupFirebase() (*firebase.App, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
