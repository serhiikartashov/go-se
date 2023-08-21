package main

// import "fmt"

// import "fmt"

func main() {
	// fmt.Println("Hi there!")

	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // first initializatio the value
	// card := newCard()
	// card = "Five of Diamonds"

	// arrays / slices
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of spades")
	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards := newDeck()
	// cards.print()
	// split deck on two parts with splitterator 5
	// deal(cards, 5)
	// fmt.Println(cards.toString())
	// cards.safeToFile("my_cards")


	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamons"
// }

