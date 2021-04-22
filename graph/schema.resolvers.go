package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bradmccoydev/graphql/graph/generated"
	"github.com/bradmccoydev/graphql/graph/model"
)

func (r *mutationResolver) CreateApplicationMetadataRegistry(ctx context.Context, input model.CreateApplicationMetadataRegistryInput) (*model.ApplicationMetadataRegistry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateApplicationMetadataRegistry(ctx context.Context, input model.UpdateApplicationMetadataRegistryInput) (*model.ApplicationMetadataRegistry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteApplicationMetadataRegistry(ctx context.Context, input model.DeleteApplicationMetadataRegistryInput) (*model.ApplicationMetadataRegistry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetApplicationMetadataRegistry(ctx context.Context, id string, version string) (*model.ApplicationMetadataRegistry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListApplicationMetadataRegistries(ctx context.Context, filter *model.TableApplicationMetadataRegistryFilterInput, limit *int, nextToken *string) (*model.ApplicationMetadataRegistryConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
