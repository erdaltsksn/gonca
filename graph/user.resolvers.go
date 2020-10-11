package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	user := model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	payload := &model.CreateUserPayload{
		ID: user.ID,
	}

	return payload, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	// Validate login input
	inputUser := &model.User{
		Email:    input.Email,
		Password: input.Password,
	}
	if err := inputUser.Validate(); err != nil {
		return nil, err
	}

	// Get the user
	var user model.User
	if err := r.DB.Where(&model.User{Email: input.Email}).First(&user).Error; err != nil {
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
	}).SignedString([]byte("gonca_auth_secret"))
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
	}).SignedString([]byte("gonca_auth_secret"))
	if err != nil {
		return nil, err
	}

	if err := r.Redis.Set(context.Background(), user.Email, refreshToken, 0).Err(); err != nil {
		return nil, err
	}

	payload := &model.LoginPayload{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return payload, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
