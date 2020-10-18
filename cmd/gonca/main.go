package main

//go:generate go run ../gqlgen/main.go

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/erdaltsksn/cui"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"

	"github.com/erdaltsksn/gonca/auth"
	"github.com/erdaltsksn/gonca/database"
	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/graph"
)

func main() {
	// Read in config file and ENV variables if set.
	dir, err := os.Getwd()
	if err != nil {
		cui.Error("Couldn't get the working directory", err)
	}

	// Search config in the working directory with name ".config" (without extension).
	viper.AddConfigPath(dir)
	viper.SetConfigName(".config")

	// Read in environment variables that match.
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:")
		cui.Info(viper.ConfigFileUsed())
	}

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

	var addr string
	if viper.IsSet("server.url") {
		addr = viper.GetString("server.url")
	}
	if viper.IsSet("server.port") {
		addr += ":" + viper.GetString("server.port")
	}
	log.Fatal(http.ListenAndServe(addr, r))
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
