package main

//go:generate go run ../gqlgen/main.go

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/erdaltsksn/gonca/auth"
	"github.com/erdaltsksn/gonca/database"
	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/graph"
)

func main() {
	// Connect the database
	db := database.PostgreSQL()

	// Migrate the schema
	db.AutoMigrate(&auth.User{})

	// Define graphql config
	cfg := generated.Config{Resolvers: &graph.Resolver{}}

	cfg.Directives.Authenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		return auth.AuthenticatedDirective(ctx, next)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Compress(5),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		addHeaders,
	)

	r.Handle("/", srv)

	log.Fatal(http.ListenAndServe(":4000", r))
}

// addHeaders adds headers to context.
func addHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		for key, val := range r.Header {
			ctx = context.WithValue(ctx, key, val[0])
		}

		// Execute the next
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
