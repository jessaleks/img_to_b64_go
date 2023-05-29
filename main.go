package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func main() {
	var base64Encoding string
	var buf bytes.Buffer

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

	// Resize the image to 64x300 pixels, blur the image, increase saturation by 10% and encode it to base64

	img, _, err := image.Decode(imgResponse.Body)

	if err != nil {
		log.Fatalf("Error decoding image: %v", err)
		os.Exit(1)
	}

	newImg := imaging.Fill(img, 64, 30, imaging.Center, imaging.Linear)
	saturatedImage := imaging.AdjustSaturation(newImg, 1.2)
	blurredImg := imaging.Blur(saturatedImage, 2.5)

	if err != nil {
		log.Fatalf("Error resizing image: %v", err)
		os.Exit(1)
	}

	// detecting the image type

	switch imgResponse.Header.Get("Content-Type") {
	case "image/jpeg":
		jpeg.Encode(&buf, blurredImg, nil)
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		png.Encode(&buf, blurredImg)
		base64Encoding += "data:image/png;base64,"
	case "image/webp":
		webp.Encode(&buf, blurredImg, nil)
		base64Encoding += "data:image/webp;base64,"
	default:
		log.Fatal("Unsupported image format")
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(buf.Bytes())

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
