package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"

	"io"
	"golang.org/x/image/draw"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Got request method %s", r.Method)
    defer r.Body.Close()

    r.ParseMultipartForm(32 << 20)
    file, _, _ := r.FormFile("data")
	result, _ := makeArt(file)

    fmt.Fprint(w, result)
}

func resize(img image.Image) image.Image {

	result := image.NewRGBA(image.Rect(0, 0, 150, 200))
	draw.ApproxBiLinear.Scale(result, result.Rect, img, img.Bounds(), draw.Over, nil)
	return result

}

func makeArt(r io.Reader) (string, error) {
	
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

func main() {
    listenAddr := ":8080"
    if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
        listenAddr = ":" + val
    }
    http.HandleFunc("/api/asciiart", helloHandler)
    fmt.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}