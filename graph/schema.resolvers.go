package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cobyforrester/serve-example/graph/generated"
	"github.com/cobyforrester/serve-example/graph/model"
	scalars "github.com/cobyforrester/serve-example/schema"
)

func (r *mutationResolver) TransformImage(ctx context.Context, input *model.ImageInstructions) (*scalars.Image, error) {
	image := input.Image
	// fmt.Println(image.Image)

	return &image, nil
}

func (r *queryResolver) Image(ctx context.Context) (*scalars.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
