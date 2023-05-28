package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/chai2010/webp"
)

func main() {
	arg := os.Args[1]
	if arg == "" || arg == "\n" {
		os.Exit(1)
	}

	imgResponse, err := http.Get(arg)
	if err != nil {
		log.Fatalf("Error fetching image: %v", err)
		os.Exit(1)
	}

	if imgResponse.StatusCode != http.StatusOK {
		log.Fatalf("Error fetching image: %v", imgResponse.Status)
		os.Exit(1)
	}

	defer imgResponse.Body.Close()

	if err != nil {
		log.Fatalf("Error reading image: %v", err)
		os.Exit(1)
	}

	// Resize the image to 64x300 pixels

	img, _, err := image.Decode(imgResponse.Body)

	newImg := image.NewRGBA(image.Rect(0, 0, 64, 30))

	draw.
	//converting image.Image to bytes
	//creates a new reader that reads from resized

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	var base64Encoding string
	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	case "image/webp":
		base64Encoding += "data:image/webp;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	fmt.Println(base64Encoding)
}

func decode(imgResponse *http.Response, err error) image.Image {
	var image image.Image
	switch imgResponse.Header.Get("Content-Type") {
	case "image/jpeg":
		image, err = jpeg.Decode(imgResponse.Body)
	case "image/png":
		image, err = png.Decode(imgResponse.Body)
	case "image/webp":
		image, err = decodeWebP(imgResponse.Body)
	default:
		log.Fatal("Unsupported image format")
	}
	if err != nil {
		log.Fatal(err)
	}
	return image
}

func decodeWebP(r io.Reader) (image.Image, error) {
	img, err := webp.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
