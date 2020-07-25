package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	fmt.Println("Connecting to: ", dbURL)
	// Make db client
	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed connecting to postgres: %v", err)
	}
	defer client.Close()

	// Run migrations
	ctx := context.Background()
	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
	//migrate.WithDropIndex(true),
	//migrate.WithDropColumn(true),

	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	graphResolver := &graph.Resolver{
		DB: client,
	}
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: graphResolver}))

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

	mux := http.NewServeMux()

	// Disable introspection and playground for production
	if !isProduction {
		srv.Use(extension.Introspection{})
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

	mux.Handle("/query", srv)

	handler := cors.Default().Handler(mux)

	log.Printf("Listening at http://localhost:%s/query", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func applyDotEnv() string {
	env := os.Getenv("FC_ENV")
	if env == "" {
		env = "development"
	}
	godotenv.Load(".env." + env + ".local")
	if env != "test" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return env
}
