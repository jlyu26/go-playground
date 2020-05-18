package main

import "fmt"

func main() {
	// card := "..."
	var card string = newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
