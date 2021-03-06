package main

func main() {
	//cards := newDeck()
	//append does not modify the slice but returns a new slice

	//cards.print()

	//No need to imports to access 'deal' func as it is in 'main' package

	//hand, remainingCards := deal(cards, 5)
	//prints cards in hand and ones remaining ; 0-4 in hand
	//hand.print()
	//remainingCards.print()

	//print list of cards as string
	//fmt.Println(cards.toString())

	//saving list of cards to hard drive
	//cards.saveToFile("my_cards")
	//file reflects in explorer as plain text file

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	//prints shuffled set of cards
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
