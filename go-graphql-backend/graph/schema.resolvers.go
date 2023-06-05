package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"go-graphql-backend/graph/model"
	"time"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID:          id,
		Description: "User " + id + " description here. (coming from the Go backend)",
	}, nil
}

// TestSubscription is the resolver for the testSubscription field.
func (r *subscriptionResolver) TestSubscription(ctx context.Context) (<-chan string, error) {
	ch := make(chan string, 1)
	go func() {
		for i := 0; i < 100; i++ {
			select {
			case ch <- fmt.Sprintf("Hello! These are generated from the go backend. (iter: %d)", i):
			default:
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return ch, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
