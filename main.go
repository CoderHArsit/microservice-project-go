package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sethvargo/go-envconfig"
)

// AppConfig holds configuration details loaded from environment variables
type AppConfig struct {
	AccountURL string `env:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `env:"CATALOG_SERVICE_URL"`
	OrderURL   string `env:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	ctx := context.Background()
	err := envconfig.Process(ctx, &cfg)
	if err != nil {
		log.Fatalf("Failed to load environment configuration: %v", err)
	}

	srv, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)
	if err != nil {
		log.Fatalf("Failed to create GraphQL server: %v", err)
	}

	http.Handle("/graphql", handler.NewDefaultServer(srv.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	log.Println("Server is running at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
