package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
)

// AuthenticatedDirective ...
func AuthenticatedDirective(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	if ctx.Value("Authorization") == nil {
		return nil, errors.New("Empty access token")
	}

	tokenString := fmt.Sprint(ctx.Value("Authorization"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprint("Unexpected signing method:", token.Header["alg"]))
		}

		return []byte(viper.GetString("auth.secret")), nil
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
