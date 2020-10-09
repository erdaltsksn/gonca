package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/model"
)

func (r *queryResolver) Ping(ctx context.Context) (*model.PingPayload, error) {
	payload := model.PingPayload{
		Message: "Pong",
	}

	return &payload, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
