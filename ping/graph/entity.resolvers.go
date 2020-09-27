package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/erdaltsksn/gonca/ping/generated"
	"github.com/erdaltsksn/gonca/ping/model"
)

func (r *entityResolver) FindPingByMessage(ctx context.Context, message string) (*model.Ping, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
