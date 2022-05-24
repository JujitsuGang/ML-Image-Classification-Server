package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//array of datasets
var datasets []Dataset

func main() {
	readConfig("./config.json")

	c.Cyan("reading images datasets")
	tStart := time.Now()
	datasets = readDataset("./dataset")
	fmt.Print("time spend reading images: ")
	fmt.Println(time.Since(tStart))
	fmt.Println("total folders scanned: " + strconv.Itoa(len(datasets[0])))

	numImages := 0
	f