package main

import (
	"fmt"
	"net/http"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func checkAndReturn(err error, w http.ResponseWriter, msg st