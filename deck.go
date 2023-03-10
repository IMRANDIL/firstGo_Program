package main

import "fmt"

//create a new type of deck
// which is slice of string


type deck []string


func newDec() deck{
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _,suits := range cardSuits{
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}



func (d deck) print(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}


func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}