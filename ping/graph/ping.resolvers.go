package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/erdaltsksn/gonca/ping/generated"
	"github.com/erdaltsksn/gonca/ping/model"
)

func (r *queryResolver) Ping(ctx context.Context) (*model.Ping, error) {
	msg := &model.Ping{
		Message: "Pong",
	}

	return msg, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
