package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/cobyforrester/image-transform/graph/model"

type Resolver struct {
	imageInstructions model.ImageInstructions
	imageJSONInput	model.ImageJSONInput
}
