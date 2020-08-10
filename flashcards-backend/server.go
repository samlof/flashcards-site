package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"flashcards-backend/ent"
	"flashcards-backend/ent/migrate"
	"flashcards-backend/graph"
	"flashcards-backend/graph/generated"
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
	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true),
		//err = client.Schema.WriteTo(ctx, os.Stdout, migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true))
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	graphResolver := &graph.Resolver{
		DB: client.Debug(),
	}
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: graphResolver}))
	configureGqlServer(srv)

	mux := http.NewServeMux()

	mux.Handle("/query", srv)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			log.Printf("Error with healthcheck: %v", err)
		}
	})

	// Enable introspection and playground for development
	if !isProduction {
		srv.Use(extension.Introspection{})
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground\n", port)
	}
	handler := cors.Default().Handler(mux)

	// Setup close handler
	httpServer := http.Server{
		Addr:              ":" + port,
		Handler:           handler,
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

func configureGqlServer(srv *handler.Server) {

	// Copied from handler.NewDefaultServer

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
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
