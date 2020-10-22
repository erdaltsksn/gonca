package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
)

// ParseJwtToken parse, validate an token string.
// Returns a `*jwt.Token`, `claims` and `error` in case of a failed validation.
func ParseJwtToken(ctx context.Context) (*jwt.Token, jwt.MapClaims, error) {
	if ctx.Value("Authorization") == nil {
		return nil, nil, errors.New("Empty access token")
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
			return nil, nil, errors.New("Token is expired")
		}

		return nil, nil, errors.New(fmt.Sprint("Failed to parse JWT token:", err))
	}

	// Get claims for token.
	claims, _ := token.Claims.(jwt.MapClaims)

	return token, claims, err
}
