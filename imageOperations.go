package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/anthonynsimon/bild/effect"
	"github.com/nfnt/resize"
)

type imgRGBA [][]float64

func dataToImage(data []byte, imageExtension string) (image.Image, error) {
	reader := bytes.NewReader(data)
	var img image.Image
	var err error
	switch imageExtension {
	case "png":
		img, err = png.Decode(reader)
	case "jpg", "jpeg":
		img, err = jpeg.Decode(reader)
	default:
		img = nil
	}
	if err != nil {
		return img, err
	}
	return img, err
}

func imageToData(img image.Image, imageExtension string) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	switch imageExtension {
	case "png":
		err = png.E