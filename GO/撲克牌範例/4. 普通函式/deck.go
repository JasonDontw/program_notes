package main

import "fmt"
//create a new type of 'deck'
//which is a slice pf strings

type deck [] string //設定一個型別名稱叫deck，型別類行為[]string的slice array

func newDeck() deck{ //設定一個fun叫newDeck且return deck
	cards:=deck{}
	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Tow","Three","Four"}

	for _, suit :=range cardSuits{ //_代表不用使用到
		for _, value :=range cardValues{
			cards=append(cards,value + " of " + suit)
		} 
	}
	return cards
}

func (d deck) printdeck(){ //設定一個fun叫print()用類別為deck的d變數接收傳入變數
	for i, card :=range d{
		fmt.Println(i,card)
	}
}

func deal(d deck,handSize int)(deck,deck){ //設定一個fun叫deal傳入2個變數且return 2個deck型態變數
	return d[:handSize],d[handSize:]
}