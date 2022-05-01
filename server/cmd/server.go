package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/purefun/todo-example/server/app/cluster"
	"github.com/purefun/todo-example/server/app/handlers"
	"github.com/purefun/todo-example/server/gql"
	"github.com/purefun/todo-example/server/gql/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cluster := cluster.NewCluster()
	handlers := handlers.NewHandlers(cluster)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &gql.Resolver{Handlers: handlers}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
