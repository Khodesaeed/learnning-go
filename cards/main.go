package main

func main() {

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards.txt")
	// cards := newDeckFromFile("my_cards.txt")
	cards := createDeck()
	cards.shuffle()
	cards.print()
}
