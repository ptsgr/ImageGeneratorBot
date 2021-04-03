package image_creator

import (
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"

	"github.com/ptsgr/ImageGeneratorBot/pkg/hex2rgb"
	"github.com/spf13/viper"
)

type Image struct {
	Image *image.RGBA
}

type ImageProperties struct {
	BackgroundImage *image.Image
	BackgraundColor string
	ImageWigth      int
	ImageHeight     int
	Text            string
	TextColor       string
}

func (imageProperties *ImageProperties) InitImageProperties() {
	imageProperties.ImageHeight = viper.GetInt("ImageProperties.ImageHeight")
	imageProperties.ImageWigth = viper.GetInt("ImageProperties.ImageWigth")
	imageProperties.BackgraundColor = viper.GetString("ImageProperties.BackgraundColor")
}

func (img *Image) CreateImage(out io.Writer) error {
	var imageProperties ImageProperties
	imageProperties.InitImageProperties()
	img.Image = image.NewRGBA(image.Rect(0, 0, imageProperties.ImageWigth, imageProperties.ImageHeight))
	clr, err := hex2rgb.ParsingHex(imageProperties.BackgraundColor)
	if err != nil {
		log.Fatalf("Error parse color from config: %s", err.Error())
	}
	draw.Draw(img.Image, img.Image.Bounds(), image.NewUniform(clr.ToRGB()), image.Point{}, draw.Src)
	return png.Encode(out, img.Image)
}
