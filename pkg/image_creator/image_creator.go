package image_creator

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"

	"github.com/spf13/viper"
)

type Image struct {
	Image *image.RGBA
}

type ImageProperties struct {
	BackgroundImage *image.Image
	BackgraundColor [3]uint8
	ImageWigth      int
	ImageHeight     int
	Text            string
	TextColor       string
}

func (imageProperties *ImageProperties) InitImageProperties() {
	imageProperties.ImageHeight = viper.GetInt("ImageProperties.ImageHeight")
	imageProperties.ImageWigth = viper.GetInt("ImageProperties.ImageWigth")
	imageProperties.BackgraundColor[0] = uint8(viper.GetInt("ImageProperties.BackgraundColor.Red"))
	imageProperties.BackgraundColor[1] = uint8(viper.GetInt("ImageProperties.BackgraundColor.Green"))
	imageProperties.BackgraundColor[2] = uint8(viper.GetInt("ImageProperties.BackgraundColor.Blue"))

}

func (img *Image) CreateImage(out io.Writer) error {
	var imageProperties ImageProperties
	imageProperties.InitImageProperties()
	img.Image = image.NewRGBA(image.Rect(0, 0, imageProperties.ImageWigth, imageProperties.ImageHeight))
	clr := color.RGBA{imageProperties.BackgraundColor[0], imageProperties.BackgraundColor[1], imageProperties.BackgraundColor[2], 255}
	draw.Draw(img.Image, img.Image.Bounds(), image.NewUniform(clr), image.Point{}, draw.Src)
	return png.Encode(out, img.Image)
}
