package helper

import (
	"image"
	"image/color"

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
	dstImage := imaging.Blur(i, level)
	return dstImage, nil
}

func GrayScale(i image.Image) (image.Image, error) {
	dstImage := imaging.Grayscale(i)
	return dstImage, nil
}

func Invert(i image.Image) (image.Image, error) {
	dstImage := imaging.Invert(i)
	return dstImage, nil
}

func Rotate(i image.Image, angle float64) (image.Image, error) {
	dstImage := imaging.Rotate(i, angle, color.RGBA{145, 23, 123, 0})
	return dstImage, nil
}
