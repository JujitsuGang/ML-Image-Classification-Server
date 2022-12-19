package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ImageModel struct {
	File []byte `json:"file"`
}

type Route struct {
	Name        string
	Method      string
	Pattern     s