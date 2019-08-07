package main 

import "fmt"
		


func main() {
	card :=newDeck()
	hand,remainingCards:=deal(card,5) //因為deal會返回2個desk type的變數所以要2個變數來接
	
	hand.printdeck() 
	fmt.Println()
	remainingCards.printdeck()
	
}
