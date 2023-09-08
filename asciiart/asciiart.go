package asciiart

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"

	"io"
	"log"
	"golang.org/x/image/draw"
)

func resize(img image.Image) image.Image {

	result := image.NewRGBA(image.Rect(0, 0, 150, 200))
	draw.ApproxBiLinear.Scale(result, result.Rect, img, img.Bounds(), draw.Over, nil)
	return result

}

func Draw(r io.Reader) (string, error) {
	
	img, format, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image format: ", format)
	fmt.Println("Original image size: ", img.Bounds())
	resized := resize(img)
	fmt.Println("Resized image: ", resized.Bounds())

	levels := []string{" ", "░", "▒", "▓", "█"}

	var b bytes.Buffer


	for y := resized.Bounds().Min.Y; y < resized.Bounds().Max.Y; y++ {
		for x := resized.Bounds().Min.X; x < resized.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			b.WriteString(levels[level])
		}
		b.WriteString("\n")
	}

	return b.String(), nil
}