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
	readConfig("./config.