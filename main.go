package main

//package main means ...it will create a executable file for us....
// if package main...then func main required...




func main(){
	cards := deck{"Ace of Diamonds",newCard()}

	cards = append(cards, "Six of Spades")

cards.print()

	
}


func newCard() string{
	return "The Ace of Spades"
}