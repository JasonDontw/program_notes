package main 

import (
	"fmt" 
//	"io/ioutil"
)


func main() {
	cards := newDeck()
	s:=[]string{"a","b"} //自行定義數值使用{}
	x:=[]string(s)		 //將其他定義好的變數塞入要用()	
	fmt.Println(x)
	fmt.Println() //////
	fmt.Println(cards.toString())
	fmt.Println() //////
	cards.saveToFile("my_cards")
}
