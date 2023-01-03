package main

import "fmt"

//package main means ...it will create a executable file for us....
// if package main...then func main required...




func main(){
	card := newCard()

	fmt.Println(card)
}


func newCard() string{
	return "The Ace of Spades"
}