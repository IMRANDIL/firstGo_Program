package main

import "fmt"

//package main means ...it will create a executable file for us....
// if package main...then func main required...




func main(){
	cards := []string{"Ace of Diamonds",newCard()}

	cards = append(cards, "Six of Spades")

	for i, card := range cards{
		fmt.Println(i,card)
	}

	
}


func newCard() string{
	return "The Ace of Spades"
}