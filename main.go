package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	_ "image/jpeg"

	"golang.org/x/image/draw"
)

func resize(img image.Image) image.Image {

	result := image.NewRGBA(image.Rect(0, 0, 150, 200))
	draw.ApproxBiLinear.Scale(result, result.Rect, img, img.Bounds(), draw.Over, nil)
	return result

}

func main() {
	
	args := os.Args
	if len(args) == 1 {
		log.Fatal("missing filename")
	}
	
	imageFile, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()

	// Consider using the general image.Decode as it can sniff and decode any registered image format.
	
	img, format, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image format: ", format)
	// fmt.Println(img)
	fmt.Println("Original image size: ", img.Bounds())
	resized := resize(img)
	fmt.Println("Resized image: ", resized.Bounds())

	levels := []string{" ", "░", "▒", "▓", "█"}


	for y := resized.Bounds().Min.Y; y < resized.Bounds().Max.Y; y++ {
		for x := resized.Bounds().Min.X; x < resized.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}