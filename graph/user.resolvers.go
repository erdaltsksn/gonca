package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/erdaltsksn/gonca/auth"
	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/generated/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	return auth.CreateUser(input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	return auth.Login(input)
}

func (r *mutationResolver) Logout(ctx context.Context) (*model.LogoutPayload, error) {
	return auth.Logout(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
