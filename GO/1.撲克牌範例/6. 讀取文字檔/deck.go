package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)


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

func (d deck)toString() string{
	return strings.Join(d,",") //將陣列裡面的字串合成一個大字串且用","隔開
}

func (d deck) saveToFile(filename string) error { 
		return ioutil.WriteFile(filename,[]byte(d.toString()),0666) //將傳進來的值當作filename
}																 	//將d的值轉成一個string且變成[]type
																 	//且默認權限為0666(最大權限)

func newDeckFromFile(path string)deck{
	bs,err := ioutil.ReadFile(path)
	if err!=nil{ //nil等於null
		fmt.Println("error:",err)
		os.Exit(1) //離開此函式
	}
	s:=strings.Split(string(bs),",") //將string(bs)切割，遇到,就切割
	return deck(s)
}