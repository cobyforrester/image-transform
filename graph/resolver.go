package graph

import "github.com/cobyforrester/image-transform/graph/model"

// go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	imageInstructions model.ImageInstructions
}
