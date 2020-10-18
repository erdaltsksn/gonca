package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/erdaltsksn/gonca/database"
	"github.com/erdaltsksn/gonca/generated/model"
)

// CreateUser ...
func CreateUser(input model.CreateUserInput) (*model.CreateUserPayload, error) {
	user := User{
		Email:    input.Email,
		Password: input.Password,
	}

	if err := database.PostgreSQL().Create(&user).Error; err != nil {
		return nil, err
	}

	payload := &model.CreateUserPayload{
		ID: user.ID,
	}

	return payload, nil
}

// Login ...
func Login(input model.LoginInput) (*model.LoginPayload, error) {
	// Validate login input
	inputUser := &User{
		Email:    input.Email,
		Password: input.Password,
	}
	if err := inputUser.Validate(); err != nil {
		return nil, err
	}

	// Get the user
	var user User
	if err := database.PostgreSQL().Where(&User{Email: input.Email}).First(&user).Error; err != nil {
		return nil, err
	}

	// Check user password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, err
	}

	// Generate tokens
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	}).SignedString([]byte(viper.GetString("auth.secret")))
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
	}).SignedString([]byte(viper.GetString("auth.secret")))
	if err != nil {
		return nil, err
	}

	if err := database.Redis().Set(context.Background(), user.Email, refreshToken, 0).Err(); err != nil {
		return nil, err
	}

	payload := &model.LoginPayload{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return payload, nil
}

// Logout ...
func Logout(ctx context.Context) (*model.LogoutPayload, error) {
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

	// Get claims for token.
	claims, _ := token.Claims.(jwt.MapClaims)

	if token.Valid && claims["sub"] != "" {
		database.Redis().Del(context.Background(), fmt.Sprint(claims["sub"]))
	} else {
		return nil, errors.New("Something went wrong on backend")
	}

	payload := &model.LogoutPayload{
		Message: "Logout successfully",
	}

	return payload, nil
}
