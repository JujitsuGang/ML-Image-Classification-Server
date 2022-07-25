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

type Dataset