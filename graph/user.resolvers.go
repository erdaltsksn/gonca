package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/model"
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
