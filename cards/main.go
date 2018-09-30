package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("*********** Shuffled **************")
	cards.shuffle()
	cards.print()
}
