package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	// go will initiate a map with zero value, which is map[]
	// var colors map[string]string

	colors := make(map[string]string)
	fmt.Println(colors)
}
