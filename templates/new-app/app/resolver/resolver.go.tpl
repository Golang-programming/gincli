package resolver

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) QueryExample(ctx context.Context) (string, error) {
	return "GraphQL Query Example", nil
}

func (r *Resolver) MutationExample(ctx context.Context) (string, error) {
	return "GraphQL Mutation Example", nil
}
