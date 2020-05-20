package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/anthonynsimon/bild/effect"
	"github.com/nfnt/resize"
)

type imgRGBA [][]