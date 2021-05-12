package helper

import (
	"image"

	"github.com/cobyforrester/image-transform/graph/model"
)

func RunTransformations(i *image.Image, input model.ImageJSONInput) (image.Image, error) {
	if input.Image.Options.Scale != nil {
		conf, err := B64ToImageConfig(input.Image.Base64)
		if err != nil {
			return nil, err
		}
		*i, _ = Scale(*i, conf, *input.Image.Options.Scale)
	}
	if input.Image.Options.Blur != nil {
		*i, _ = Blur(*i, *input.Image.Options.Blur)
	}
	if input.Image.Options.Rotate != nil {
		*i, _ = Rotate(*i, *input.Image.Options.Rotate)
	}
	if input.Image.Options.Grayscale != nil && *input.Image.Options.Grayscale == true {
		*i, _ = GrayScale(*i)
	}
	if input.Image.Options.Invert != nil && *input.Image.Options.Invert == true {
		*i, _ = Invert(*i)
	}
	return *i, nil
}
