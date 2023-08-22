package text2img

import (
	"image/color"
	"math/rand"
)

// Color contains a good conbination of backgroundColor and textColor
type Color struct {
	BgColor color.RGBA
	FgColor color.RGBA
}

type palette struct {
	BgColor string
	FgColor string
}

var palettes = []palette{
	{"#003d47", "#fff"},
	{"#128277", "#fff"},
	{"#d24136", "#fff"},
	{"#eb8a3e", "#fff"},
	{"#ebb582", "#fff"},
	{"#785a46", "#fff"},
	{"#bc6d4f", "#fff"},
	{"#1e1f26", "#fff"},
	{"#283655", "#fff"},
	{"#4d648d", "#fff"},
	{"#265c00", "#fff"},
	{"#faaf08", "#fff"},
	{"#fa812f", "#fff"},
	{"#fa4032", "#fff"},
	{"#6c5f5b", "#fff"},
	{"#cdab81", "#fff"},
	{"#4f4a45", "#fff"},
	{"#04202c", "#fff"},
	{"#304040", "#fff"},
	{"#5b7065", "#fff"},
	{"#1e0000", "#fff"},
	{"#500805", "#fff"},
	{"#9d331f", "#fff"},
	{"#68a225", "#fff"},
	{"#2c4a52", "#fff"},
	{"#537072", "#fff"},
	{"#8e9b97", "#fff"},
	{"#d8412f", "#fff"},
	{"#fe7a47", "#fff"},
	{"#867666", "#fff"},
	{"#e1b80d", "#fff"},
	{"#003b46", "#fff"},
	{"#07575b", "#fff"},
	{"#66a5ad", "#fff"},
	{"#af6c59", "#fff"},
	{"#e68f71", "#fff"},
	{"#021c1e", "#fff"},
	{"#004445", "#fff"},
	{"#2c7873", "#fff"},
	{"#6fb98f", "#fff"},
	{"#434343", "#fff"},
	{"#767676", "#fff"},
	{"#c16707", "#fff"},
	{"#f08d16", "#fff"},
	{"#77262a", "#fff"},
	{"#9e2d29", "#fff"},
	{"#c35d44", "#fff"},
	{"#202d35", "#fff"},
	{"#0e3c54", "#fff"},
	{"#2a677c", "#fff"},
	{"#4f3538", "#fff"},
	{"#66443b", "#fff"},
	{"#c29f83", "#fff"},
	{"#210e3b", "#fff"},
	{"#4b194c", "#fff"},
	{"#872b76", "#fff"},
	{"#fdffff", "#333"},
	{"#fcfdfe", "#333"},
	{"#f4ebdb", "#333"},
}

// PickColor picks a color
func PickColor() Color {
	n := rand.Intn(len(palettes) + 1)
	bgStr, fgStr := palettes[n].BgColor, palettes[n].FgColor

	// FIXME manage error
	bg, _ := hexToColor(bgStr)
	fg, _ := hexToColor(fgStr)

	return Color{BgColor: bg, FgColor: fg}
}
