package main

import "fmt"

func main() {

	cards := deck{"some string", "some string"}

	for _, card := range cards {
		fmt.Println(card)
	}
}
