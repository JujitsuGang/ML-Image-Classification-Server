
package main

import (
	"fmt"
	"image"
	"sort"
)

type Neighbour struct {
	Dist  float64
	Label string
}

func euclideanDist(img1, img2 [][]float64) float64 {