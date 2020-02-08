package graphql_go_test

import (
	"context"

	"github.com/maiconcarraro/graphql-go-test/models" // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTask(ctx context.Context, input NewTask) (*models.Task, error) {
	return createMockTask(input)
}

type taskResolver struct{ *Resolver }

func (m *taskResolver) User(ctx context.Context, obj *models.Task) (*models.User, error) {
	return findMockUser(obj.UserID)
}

func (r *Resolver) Task() TaskResolver {
	return &taskResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Tasks(ctx context.Context) ([]*models.Task, error) {
	return getMockTasks(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return getMockUsers(), nil
}
