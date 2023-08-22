package text2img

import (
	"image/jpeg"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	path := "fonts/MPLUS1Code-Bold.ttf"
	d, err := NewDrawer(Params{
		FontPath: path,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	img, err := d.Draw("text2img generates the image from a text")
	if err != nil {
		t.Fatal(err.Error())
	}
	file, err := os.Create("test.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()
	if err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
		t.Fatal(err.Error())
	}
}
