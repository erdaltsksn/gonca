package main

//go:generate go run ../gqlgen/main.go

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/graph"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", srv)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
