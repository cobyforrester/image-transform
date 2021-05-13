package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cobyforrester/image-transform/graph/generated"
	"github.com/cobyforrester/image-transform/graph/model"
	"github.com/cobyforrester/image-transform/helper"
	scalars "github.com/cobyforrester/image-transform/schema"
)

func (r *mutationResolver) TransformImage(ctx context.Context, input model.ImageInstructions) (*scalars.Image, error) {
	image := input.Image
	// fmt.Println(image.Image)

	return &image, nil
}

func (r *mutationResolver) TransformJSONImage(ctx context.Context, input model.ImageJSONInput) (string, error) {
	img, err := helper.B64ToImage(input.Image.Base64)
	if err != nil {
		return "", err
	}
	_, err = helper.RunTransformations(&img, input)
	if err != nil {
		return "", err
	}

	encodedStr, err := helper.ImageToB64(img)
	if err != nil {
		return "", err
	}
	return encodedStr, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
