package image_creator

import (
	"image"
	"image/png"
	"io"
)

const (
	wigth  = 360
	height = 100
)

func CreateImage(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, wigth, height))
	png.Encode(out, img)
}
