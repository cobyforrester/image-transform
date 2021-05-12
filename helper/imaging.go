package helper

import (
	"image"

	"github.com/disintegration/imaging"
)

func Scale(i image.Image, config image.Config, scale float64) (image.Image, error) {
	// height := int(float64(config.Width) * scale)

	// setting height to 0 auto scales it to match width dimensions
	width := int(float64(config.Height) * scale)
	resizedImg := imaging.Resize(i, width, 0, imaging.Lanczos)
	return resizedImg, nil
}

func Blur(i image.Image, level float64) (image.Image, error) {
	// height := int(float64(config.Width) * scale)

	// setting height to 0 auto scales it to match width dimensions
	dstImage := imaging.Blur(i, level)
	return dstImage, nil
}
