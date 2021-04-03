package hex2rgb

import (
	"errors"
	"fmt"
	"image/color"
	"regexp"
	"strings"
)

const (
	hexRegexString = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	hexFormat      = "#%02x%02x%02x"
	hexShortFormat = "#%1x%1x%1x"
	hexToRGBFactor = 17
)

var (
	ErrBadColor = errors.New("Parsing of color failed, Bad Color")
	hexRegex    = regexp.MustCompile(hexRegexString)
)

type Hex struct {
	hex string
}

func ParsingHex(s string) (*Hex, error) {
	s = strings.ToLower(s)

	if !hexRegex.MatchString(s) {
		return nil, ErrBadColor
	}
	return &Hex{hex: s}, nil
}

func (h *Hex) ToRGB() *color.RGBA {
	clr := new(color.RGBA)
	if len(h.hex) == 4 {
		fmt.Sscanf(h.hex, hexShortFormat, &clr.R, &clr.G, &clr.B)
		clr.R *= hexToRGBFactor
		clr.G *= hexToRGBFactor
		clr.B *= hexToRGBFactor
	} else {
		fmt.Sscanf(h.hex, hexFormat, &clr.R, &clr.G, &clr.B)
	}
	clr.A = 255
	return clr
}
