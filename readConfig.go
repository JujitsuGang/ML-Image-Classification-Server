package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerIP   string   `json:"serverIP"`
	ServerPort string   `json:"serverPort"`
	ImgWidth   int      `json:"imgWidth"`
	ImgHeigh   