package main

import "fmt"

func main() {
	// These are the three ways that we can create map.
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#FFFFFF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex of the", color, "color, is", hex)
	}
}
