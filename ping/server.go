package main

//go:generate go run ../cmd/gqlgen/main.go

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/graphql/handler/debug"

	"github.com/erdaltsksn/gonca/ping/generated"
	"github.com/erdaltsksn/gonca/ping/graph"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	// srv.Use(&debug.Tracer{})

	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":4001", nil))
}
