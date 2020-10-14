package main

//go:generate go run ../gqlgen/main.go

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/erdaltsksn/gonca/database"
	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/graph"
	"github.com/erdaltsksn/gonca/model"
)

func main() {
	// Connect the database
	db := database.PostgreSQL()

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Define graphql config
	cfg := generated.Config{Resolvers: &graph.Resolver{}}

	cfg.Directives.Authenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		if ctx.Value("Authorization") == nil {
			return nil, errors.New("Empty access token")
		}

		tokenString := fmt.Sprint(ctx.Value("Authorization"))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(fmt.Sprint("Unexpected signing method:", token.Header["alg"]))
			}

			return []byte("gonca_auth_secret"), nil
		})
		if err != nil {
			var errExpired *jwt.TokenExpiredError
			if errors.As(err, &errExpired) {
				return nil, errors.New("Token is expired")
			}

			return nil, errors.New(fmt.Sprint("Failed to parse JWT token:", err))
		}

		if token.Valid {
			return next(ctx)
		}

		return nil, errors.New("Unauthorized, you are not allowed to access")
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
