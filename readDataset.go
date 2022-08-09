package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strconv"
	"strings"
)

//each image is [][]float64, is a array of pixels
type ImgDataset [][][]float64

type Dataset map[string]ImgDataset

func byteArrayToFloat64Array(b []byte) []float64 {
	var f []float64
	for i := 0; i < len(b); i++ {
		val, _ := strconv.ParseFloat(string(b[i]), 64)
		f = append(f, val)
	}
	return f
}

func readImage(path string) image.Image {
	//open image file
	dat, err := ioutil.ReadFile(path)
	check(err)

	pathSplited := strings.Split(path, "."