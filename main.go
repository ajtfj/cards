package main

const FILE = "my_cards"

func main() {
	cards := newDeckFromFile(FILE)
	cards.shuffle()
	hand, rest := deal(cards, 5)
	cards.print()
	hand.print()
	rest.print()
}
