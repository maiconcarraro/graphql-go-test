package graphql_go_test

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTask(ctx context.Context, input NewTask) (*Task, error) {
	return createMockTask(input)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Tasks(ctx context.Context) ([]*Task, error) {
	return getMockTasks(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return getMockUsers(), nil
}
