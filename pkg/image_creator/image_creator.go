package image_creator

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"strconv"

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
	BackgraundColor *hex2rgb.Hex
	ImageWigth      int
	ImageHeight     int
	Text            string
	TextColor       *hex2rgb.Hex
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
func setColorProperties(configKey, defaultValue string) *hex2rgb.Hex {
	stringColor := viper.GetString(configKey)
	clr, err := hex2rgb.ParsingHex(stringColor)
	if err != nil {
		log.Printf("Error cannot use color from config: %s", err.Error())
		clr, _ = hex2rgb.ParsingHex(defaultValue)
	}
	return clr

}

func (imageProperties *ImageProperties) InitImageProperties(keys map[string][]string) {

	var err error
	if keys["height"] != nil {
		imageProperties.ImageHeight, err = strconv.Atoi(keys["height"][0])
		if err != nil {
			log.Printf("Error parse Image Height: %s\nUse value from config(if exist) or default value.", err.Error())
			imageProperties.ImageHeight = setIntProperties("ImageProperties.ImageHeight", DefaultImageHeight)
		}
	} else {
		imageProperties.ImageHeight = setIntProperties("ImageProperties.ImageHeight", DefaultImageHeight)
	}
	if keys["wigth"] != nil {
		imageProperties.ImageWigth, err = strconv.Atoi(keys["wigth"][0])
		if err != nil {
			log.Printf("Error parse Image Wigth: %s\nUse value from config(if exist) or default value.", err.Error())
			imageProperties.ImageWigth = setIntProperties("ImageProperties.ImageWigth", DefaultImageWigth)
		}
	} else {
		imageProperties.ImageWigth = setIntProperties("ImageProperties.ImageWigth", DefaultImageWigth)
	}

	if keys["label"] != nil {
		imageProperties.Text = keys["label"][0]
	} else {
		imageProperties.Text = setStirngProperties("ImageProperties.Text", DefaultText)
	}
	if keys["color"] != nil {
		imageProperties.BackgraundColor, err = hex2rgb.ParsingHex("#" + keys["color"][0])
		if err != nil {
			log.Printf("Error parse %s backgraund color: %s", keys["color"][0], err.Error())
			imageProperties.BackgraundColor = setColorProperties("ImageProperties.BackgraundColor", DefaultBackgraundColor)
		}
	} else {
		imageProperties.BackgraundColor = setColorProperties("ImageProperties.BackgraundColor", DefaultBackgraundColor)
	}

	if keys["labelColor"] != nil {
		imageProperties.TextColor, err = hex2rgb.ParsingHex("#" + keys["labelColor"][0])
		if err != nil {
			log.Printf("Error parse %s label color: %s", keys["labelColor"][0], err.Error())
			imageProperties.TextColor = setColorProperties("ImageProperties.TextColor", DefaultTextColor)
		}
	} else {
		imageProperties.TextColor = setColorProperties("ImageProperties.TextColor", DefaultTextColor)
	}
	imageProperties.LabelFontFile = setStirngProperties("ImageProperties.labelFontFile", DefaultFontFile)
}

func (img *Image) CreateImage() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	img.Image = image.NewRGBA(image.Rect(0, 0, img.Properties.ImageWigth, img.Properties.ImageHeight))

	draw.Draw(img.Image, img.Image.Bounds(), image.NewUniform(img.Properties.BackgraundColor.ToRGB()), image.Point{}, draw.Src)
	img.AddText()

	if err := png.Encode(buffer, img.Image); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (img *Image) AddText() {

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
		Src: image.NewUniform(img.Properties.TextColor.ToRGB()),
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
