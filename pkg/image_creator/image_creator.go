package image_creator

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/ptsgr/ImageGeneratorBot/pkg/hex2rgb"
	"github.com/spf13/viper"
	"golang.org/x/image/font"
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
	LabelFontFile   string
}

const (
	DefaultBackgraundColor = "#FFF"
	DefaultImageWigth      = 640
	DefaultImageHeight     = 480
	DefaultText            = "Image Generator"
	DefaultTextColor       = "#000"
	DefaultFontFile        = "assets/Raleway-Bold.ttf"
)

var imageProperties ImageProperties

func setIntProperties(configKey string, defaultValue int) int {
	intParam := viper.GetInt(configKey)
	if intParam <= 0 {
		intParam = defaultValue
	}
	return intParam
}

func setStirngProperties(configKey, defaultValue string) string {
	stringParam := viper.GetString(configKey)
	if stringParam == "" {
		return defaultValue
	}
	return stringParam
}

func (imageProperties *ImageProperties) InitImageProperties(keys map[string][]string) {
	imageProperties.ImageHeight = setIntProperties("ImageProperties.ImageHeight", DefaultImageHeight)
	imageProperties.ImageWigth = setIntProperties("ImageProperties.ImageWigth", DefaultImageWigth)
	imageProperties.BackgraundColor = setStirngProperties("ImageProperties.BackgraundColor", DefaultBackgraundColor)
	imageProperties.Text = setStirngProperties("ImageProperties.Text", DefaultText)
	imageProperties.TextColor = setStirngProperties("ImageProperties.TextColor", DefaultTextColor)
	imageProperties.LabelFontFile = setStirngProperties("ImageProperties.labelFontFile", DefaultFontFile)
}

func (img *Image) CreateImage() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	img.Image = image.NewRGBA(image.Rect(0, 0, img.Properties.ImageWigth, img.Properties.ImageHeight))
	clr, err := hex2rgb.ParsingHex(img.Properties.BackgraundColor)
	if err != nil {
		log.Printf("Error parse color from config: %s", err.Error())
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
		log.Printf("Error parse color from config: %s", err.Error())
	}

	fontBytes, err := ioutil.ReadFile(img.Properties.LabelFontFile)
	if err != nil {
		log.Printf("Error open font file: %s", err.Error())
		return
	}

	labelFont, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Printf("Error parse fontBytes: %s", err.Error())
	}

	d := &font.Drawer{
		Dst: img.Image,
		Src: image.NewUniform(textClr.ToRGB()),
		Face: truetype.NewFace(labelFont, &truetype.Options{
			Size: float64(img.Properties.ImageHeight / 2 / 10),
		}),
	}

	d.Dot = fixed.Point26_6{
		X: (fixed.I(img.Properties.ImageWigth) - d.MeasureString(img.Properties.Text)) / 2,
		Y: fixed.I((img.Properties.ImageHeight) / 2),
	}

	d.DrawString(img.Properties.Text)
}
