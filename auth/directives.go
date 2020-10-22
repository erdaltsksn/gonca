package auth

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

// AuthenticatedDirective ...
func AuthenticatedDirective(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	token, _, err := ParseJwtToken(ctx)
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return next(ctx)
	}

	return nil, errors.New("Unauthorized, you are not allowed to access")
}
