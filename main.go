package main

func main() {
	//var card string = "Ace of Spades"
	//:= for 1st init
	//= for later init
	// cards := newDeck()

	// //cards.print()
	// hand, remainingCards := deal(cards, 5)

	// print("On hands\n\n")
	// hand.print()
	// print("Remaining cards\n\n")
	// remainingCards.print()

	// Create File
	// cards := newDeck()
	// //fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// Open File
	// cards := newDeckFromFile("my")
	// cards.print()

	//Shuffle card
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
