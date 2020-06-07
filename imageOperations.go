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
	var img imag