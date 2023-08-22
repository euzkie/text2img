package text2img

import (
	"fmt"
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

// hexToColor parses a string that represents the hex notation for a color,
// using a '#' followed by either 3 or 6 digits, and returns a color.RGBA object.
// A non-nil error means that the string doesn't rapresent a valid hex color.
// See https://stackoverflow.com/a/54200713
func hexToColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7: // Format #123456
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4: // Format #123 == #112233
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Multiplying a single hex digit by 17 is a shortcut to doubling the digit and
		// converting the resulting number to dec. It works because of how hex->dec
		// conversions work:
		// 25_16 = 2*16^1 + 5*16^0 = 32 + 5 = 37
		// When the two hex digits are the same:
		// 22_16 = 2*16^1 + 2*16^0 = 34 = 2*17
		c.R *= 17
		c.B *= 17
		c.G *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4") // TODO make the error message more clear.
	}
	return
}

// This is the old, dumb version of calcFontSize.
// func (d *drawer) calcFontSize(text string) (fontSize float64) {
// 	const padding = 4 // Not used
// 	fontSizes := []float64{128, 64, 48, 32, 24, 18, 16, 14, 12}
// 	for _, fontSize = range fontSizes {
// 		textWidth := d.calcTextWidth(fontSize, text)
// 		if textWidth < d.Width {
// 			return
// 		}
// 	}
// 	return
// }

// calcFontSize
func (d *drawer) calcFontSize(s string) float64 {
	ratio := 1.2 // See https://designcode.io/typographic-scales
	fontSize := 1.0
	for i := 0; float64(d.calcTextWidth(fontSize, s)) < (float64(d.Width) * 0.85); i++ {
		fontSize *= ratio
	}
	return fontSize
}

// TODO fix this.
func (d *drawer) calcTextWidth(fontSize float64, text string) (textWidth int) {
	var face font.Face
	if d.Font != nil {
		opts := truetype.Options{}
		opts.Size = fontSize
		face = truetype.NewFace(d.Font, &opts)
	} else {
		face = basicfont.Face7x13
	}
	for _, x := range text {
		awidth, ok := face.GlyphAdvance(rune(x))
		if ok != true {
			return
		}
		textWidth += int(float64(awidth) / 64)
	}
	return
}
