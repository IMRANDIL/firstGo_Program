package main

//package main means ...it will create a executable file for us....
// if package main...then func main required...




func main(){
	cards := newDec()

hand, remainingCard := deal(cards,5)

hand.print()
remainingCard.print()

	
}


