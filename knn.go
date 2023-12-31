
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
		neighbours[i] = temp[i]
	}

	return neighbours
}

func getMapKey(dataset map[string]ImgDataset) string {
	for k, _ := range dataset {
		return k
	}
	return ""
}

type LabelCount struct {
	Label string
	Count int
}

func averageLabel(neighbours []Neighbour) string {
	labels := make(map[string]int)
	for _, n := range neighbours {
		labels[n.Label]++
	}
	//create array from map
	var a []LabelCount
	for k, v := range labels {
		a = append(a, LabelCount{k, v})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Count > a[j].Count
	})
	fmt.Println(a)
	//send the most appeared neighbour in k
	return a[0].Label
}

func distNeighboursFromDataset(dataset Dataset, neighbours []Neighbour, input [][]float64) []Neighbour {
	//check the complete dataset, checking if each entry is a k nearest neighbour
	for l, v := range dataset {
		for i := 0; i < len(v); i++ {
			dNew := euclideanDist(v[i], input)
			neighbours = isNeighbour(neighbours, dNew, l)
		}
	}
	return neighbours
}
func knn(datasets []Dataset, imgInput image.Image) string {
	k := 6
	var neighbours []Neighbour
	var neighboursED []Neighbour
	/*
		var neighboursG []Neighbour
	*/

	imgED := EdgeDetection(imgInput)
	/*
		imgG := Grayscale(imgInput)
	*/

	histogram := imageToHistogram(imgInput)
	histogramED := imageToHistogram(imgED)
	/*
		histogramG := imageToHistogram(imgG)
	*/

	//get a key from map dataset, the key is a label
	label := getMapKey(datasets[0])
	//fill the first k neighbours
	for i := 0; i < k; i++ {
		neighbours = append(neighbours, Neighbour{euclideanDist(datasets[0][label][0], histogram), label})
		neighboursED = append(neighboursED, Neighbour{euclideanDist(datasets[1][label][0], histogramED), label})
		/*
			neighboursG = append(neighboursG, Neighbour{euclideanDist(datasets[2][label][0], histogramG), label})
		*/
	}

	neighbours = distNeighboursFromDataset(datasets[0], neighbours, histogram)
	neighboursED = distNeighboursFromDataset(datasets[1], neighboursED, histogramED)
	/*
		neighboursG = distNeighboursFromDataset(datasets[2], neighboursG, histogramG)
	*/

	neighbours = append(neighbours, neighboursED...)
	/*
		neighbours = append(neighbours, neighboursG...)
	*/

	for i := 0; i < len(neighbours); i++ {
		fmt.Print(neighbours[i].Label + " - ")
		fmt.Println(neighbours[i].Dist)
	}
	//from the k nearest neighbours, get the more frequent neighbour
	r := averageLabel(neighbours)
	return r
}