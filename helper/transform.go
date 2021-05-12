package helper

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/disintegration/imaging"
)

type MyImage struct {
	value *image.RGBA
}

func (i *MyImage) Set(x, y int, c color.Color) {
	i.value.Set(x, y, c)
}
func (i *MyImage) ColorModel() color.Model {
	return i.value.ColorModel()
}
func (i *MyImage) Bounds() image.Rectangle {
	return i.value.Bounds()
}
func (i *MyImage) At(x, y int) color.Color {
	return i.value.At(x, y)
}

func (bgImg *MyImage) drawRaw(innerImg image.Image, sp image.Point, width int, height int) {
	resizedImg := imaging.Resize(innerImg, width, height, imaging.Lanczos)
	b64, err := ImageToB64(resizedImg)
	if err != nil {
		panic("Bad Conversion")
	}
	config, err := B64ToImageConfig(b64)
	if err != nil {
		panic("Bad Conversion")
	}
	w := int(config.Width)
	h := int(config.Height)
	draw.Draw(bgImg, image.Rectangle{sp, image.Point{sp.X + w, sp.Y + h}}, resizedImg, image.ZP, draw.Src)
}
