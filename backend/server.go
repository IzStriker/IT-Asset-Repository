package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IzStriker/IT-Asset-Repository/backend/graph"
	"github.com/IzStriker/IT-Asset-Repository/backend/graph/generated"
	"github.com/IzStriker/IT-Asset-Repository/backend/neo4j"
)

const defaultPort = "8080"

func main() {
	port := getEnv("PORT", defaultPort)
	assetRepo := neo4j.AssetRepo{
		Uri:      getEnv("NEO4J_URI", ""),
		Username: getEnv("NEO4J_USERNAME", ""),
		Password: getEnv("NEO4J_PASSWORD", ""),
	}

	if err := assetRepo.Initialise(); err != nil {
		panic(err)
	}
	log.Println("Database connection successful")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	if defaultValue != "" {
		return defaultValue
	}

	panic(fmt.Sprintf("%s not set", key))
}
