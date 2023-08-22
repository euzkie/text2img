package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/Iwark/text2img"
)

var (
	fontPath  = flag.String("fontpath", "", "path to the font ttf file")
	bgImgPath = flag.String("bgimg", "", "path to the background image")
	output    = flag.String("output", "image.png", "path to the output image")
	text      = flag.String("text", "", "text to draw")
)

func main() {
	flag.Parse()

	d, err := text2img.NewDrawer(text2img.Params{
		FontPath:  *fontPath,
		BgImgPath: *bgImgPath,
	})
	if err != nil {
		log.Println(err)
		return
	}

	img, err := d.Draw(*text)
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.Create(*output)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// if err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
	if err = png.Encode(file, img); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Success")
}
