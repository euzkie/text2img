package text2img

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type drawer struct {
	BgImg             image.Image
	BgColor           *image.Uniform
	TextColor         *image.Uniform
	TextPosVertical   int
	TextPosHorizontal int
	Width             int // ImgWidth
	Height            int // ImgHeight
	Font              *truetype.Font
	FontSize          float64
	autoFontSize      bool // Remove this option: font size will always be automatic.
}

// Drawer is the main interface for this package.
type Drawer interface {
	Draw(string) (*image.RGBA, error)
	SetColors(color.RGBA, color.RGBA)
	SetTextPos(int, int)
	SetFontPath(string) error
	SetFontSize(float64)
	SetSize(int, int)
}

// Params holds the configuration for Drawer.
type Params struct {
	BgImgPath         string
	BgColor           color.RGBA
	TextColor         color.RGBA
	TextPosVertical   int
	TextPosHorizontal int
	Width             int
	Height            int
	FontPath          string
	FontSize          float64
}

// NewDrawer is a builder for Drawer.
func NewDrawer(params Params) (Drawer, error) {
	d := &drawer{}
	if params.FontPath != "" {
		err := d.SetFontPath(params.FontPath)
		if err != nil {
			return d, err
		}
	}
	if params.BgImgPath != "" {
		err := d.SetBgImg(params.BgImgPath)
		if err != nil {
			return d, err
		}
		d.SetSize(d.BgImg.Bounds().Size().X, d.BgImg.Bounds().Size().Y)
	} else {
		d.SetSize(params.Width, params.Height)
	}

	d.SetColors(params.TextColor, params.BgColor)
	d.SetFontSize(params.FontSize)

	return d, nil
}

// Draw draws the text on the image.
func (d *drawer) Draw(text string) (img *image.RGBA, err error) {
	if d.BgImg != nil {
		imgRect := image.Rectangle{image.Pt(0, 0), d.BgImg.Bounds().Size()}
		img = image.NewRGBA(imgRect)
		draw.Draw(img, img.Bounds(), d.BgImg, image.ZP, draw.Src)
	} else {
		img = image.NewRGBA(image.Rect(0, 0, d.Width, d.Height))
		draw.Draw(img, img.Bounds(), d.BgColor, image.ZP, draw.Src)
	}
	if d.autoFontSize {
		d.FontSize = d.calcFontSize(text)
	}
	textWidth := d.calcTextWidth(d.FontSize, text)

	if d.Font != nil {
		c := freetype.NewContext()
		c.SetDPI(72)
		c.SetFont(d.Font)
		c.SetFontSize(d.FontSize)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(d.TextColor)
		c.SetHinting(font.HintingNone)

		textHeight := int(c.PointToFixed(d.FontSize) >> 6)
		pt := freetype.Pt((d.Width-textWidth)/2+d.TextPosHorizontal, (d.Height+textHeight)/2+d.TextPosVertical)
		_, err = c.DrawString(text, pt)
		return
	}
	err = errors.New("Font must be specified")
	// point := fixed.Point26_6{640, 960}
	// fd := &font.Drawer{
	// 	Dst:  img,
	// 	Src:  d.TextColor,
	// 	Face: basicfont.Face7x13,
	// 	Dot:  point,
	// }
	// fd.DrawString(text)
	return
}

// SetBgImg sets the specific background image
func (d *drawer) SetBgImg(imagePath string) (err error) {
	src, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return
	}
	d.BgImg = img
	return
}

// SetColors sets the backgroundColor and the textColor.
func (d *drawer) SetColors(textColor, bgColor color.RGBA) {
	r1, g1, b1, a1 := bgColor.RGBA()
	r2, g2, b2, a2 := textColor.RGBA()
	if r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2 {
		color := PickColor()
		d.TextColor = image.NewUniform(color.FgColor)
		d.BgColor = image.NewUniform(color.BgColor)
		return
	}
	d.TextColor = image.NewUniform(textColor)
	d.BgColor = image.NewUniform(bgColor)
}

// SetFontPath sets the path to the font.
func (d *drawer) SetFontPath(fontPath string) (err error) {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return
	}
	d.Font = f
	return
}

// SetFontSize sets the size of the font.
func (d *drawer) SetFontSize(fontSize float64) {
	if fontSize > 0 {
		d.autoFontSize = false
		d.FontSize = fontSize
		return
	}
	d.autoFontSize = true
}

// SetTextPos sets the position of the text.
func (d *drawer) SetTextPos(textPosVertical, textPosHorizontal int) {
	d.TextPosVertical = textPosVertical
	d.TextPosHorizontal = textPosHorizontal
}

// SetSize sets the size. TODO take width and height from flags.
func (d *drawer) SetSize(width, height int) {
	if width <= 0 {
		d.Width = 1200
	} else {
		d.Width = width
	}
	if height <= 0 {
		d.Height = 630
	} else {
		d.Height = height
	}
}
