package image_creator

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"log"

	"github.com/ptsgr/ImageGeneratorBot/pkg/hex2rgb"
	"github.com/spf13/viper"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Image struct {
	Image      *image.RGBA
	Properties ImageProperties
}

type ImageProperties struct {
	BackgroundImage *image.Image
	BackgraundColor string
	ImageWigth      int
	ImageHeight     int
	Text            string
	TextColor       string
}

const (
	DefaultBackgraundColor = "#FFF"
	DefaultImageWigth      = 640
	DefaultImageHeight     = 480
	DefaultText            = "Image Generator"
	DefaultTextColor       = "#000"
)

var imageProperties ImageProperties

func (imageProperties *ImageProperties) InitImageProperties() {
	imageProperties.ImageHeight = viper.GetInt("ImageProperties.ImageHeight")
	if imageProperties.ImageHeight <= 0 {
		imageProperties.ImageHeight = DefaultImageHeight
	}

	imageProperties.ImageWigth = viper.GetInt("ImageProperties.ImageWigth")
	if imageProperties.ImageWigth <= 0 {
		imageProperties.ImageWigth = DefaultImageWigth
	}

	imageProperties.BackgraundColor = viper.GetString("ImageProperties.BackgraundColor")
	if imageProperties.BackgraundColor == "" {
		imageProperties.BackgraundColor = DefaultBackgraundColor
	}

	imageProperties.Text = viper.GetString("ImageProperties.Text")
	if imageProperties.Text == "" {
		imageProperties.Text = DefaultText
	}

	imageProperties.TextColor = viper.GetString("ImageProperties.TextColor")
	if imageProperties.TextColor == "" {
		imageProperties.TextColor = DefaultTextColor
	}

}

func (img *Image) CreateImage() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	img.Properties.InitImageProperties()
	img.Image = image.NewRGBA(image.Rect(0, 0, img.Properties.ImageWigth, img.Properties.ImageHeight))
	clr, err := hex2rgb.ParsingHex(img.Properties.BackgraundColor)
	if err != nil {
		log.Fatalf("Error parse color from config: %s", err.Error())
	}
	draw.Draw(img.Image, img.Image.Bounds(), image.NewUniform(clr.ToRGB()), image.Point{}, draw.Src)
	img.AddText()

	if err := png.Encode(buffer, img.Image); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (img *Image) AddText() {

	textClr, err := hex2rgb.ParsingHex(img.Properties.TextColor)
	if err != nil {
		log.Fatalf("Error parse color from config: %s", err.Error())
	}

	d := &font.Drawer{
		Dst:  img.Image,
		Src:  image.NewUniform(textClr.ToRGB()),
		Face: basicfont.Face7x13,
	}

	d.Dot = fixed.Point26_6{
		X: (fixed.I(img.Properties.ImageWigth) - d.MeasureString(img.Properties.Text)) / 2,
		Y: fixed.I((img.Properties.ImageHeight) / 2),
	}

	d.DrawString(img.Properties.Text)
}
