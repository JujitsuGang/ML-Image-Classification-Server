
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
	var dist float64
	for i := 0; i < len(img1); i++ {
		for j := 0; j < len(img1[i]); j++ {
			dist += (img1[i][j] - img2[i][j]) * (img1[i][j] - img2[i][j])
		}
	}

	return dist
}

func isNeighbour(neighbours []Neighbour, dist float64, label string) []Neighbour {
	var temp []Neighbour

	for i := 0; i < len(neighbours); i++ {
		temp = append(temp, neighbours[i])
	}
	ntemp := Neighbour{dist, label}
	temp = append(temp, ntemp)

	//now, sort the temp array
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].Dist < temp[j].Dist
	})

	for i := 0; i < len(neighbours); i++ {