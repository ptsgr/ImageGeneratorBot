package image_creator

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"

	"github.com/ptsgr/ImageGeneratorBot/pkg/hex2rgb"
	"github.com/spf13/viper"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
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

func (img *Image) CreateImage(out io.Writer) error {
	imageProperties.InitImageProperties()
	img.Image = image.NewRGBA(image.Rect(0, 0, imageProperties.ImageWigth, imageProperties.ImageHeight))
	clr, err := hex2rgb.ParsingHex(imageProperties.BackgraundColor)
	if err != nil {
		log.Fatalf("Error parse color from config: %s", err.Error())
	}
	draw.Draw(img.Image, img.Image.Bounds(), image.NewUniform(clr.ToRGB()), image.Point{}, draw.Src)
	textClr, err := hex2rgb.ParsingHex(imageProperties.TextColor)
	if err != nil {
		log.Fatalf("Error parse color from config: %s", err.Error())
	}

	img.AddText(imageProperties.ImageWigth/2, imageProperties.ImageHeight/2, imageProperties.Text, textClr.ToRGB())
	return png.Encode(out, img.Image)
}

func (img *Image) AddText(x, y int, text string, clr color.Color) {
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img.Image,
		Src:  image.NewUniform(clr),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}
