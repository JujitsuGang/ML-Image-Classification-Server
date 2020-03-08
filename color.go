
package main

import "fmt"

//Color struct, defines the color
type Color struct{}

var c Color

//DarkGray color
func (c Color) DarkGray(t string) {
	fmt.Print("\x1b[30;1m") //dark gray
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}

//Red color
func (c Color) Red(t string) {
	fmt.Print("\x1b[31;1m") //red
	fmt.Println(t)
	fmt.Print("\x1b[0m") //defaultColor
}
