package main 

func main() { 
	
	cards := deck{"a","b"} //deck型別定義在deck.go

	cards.printdeck()   //函式定義於deck.go，因為printdeck函數需要傳入deck的變數，cards是deck的type所以可以使用
	
	cardss:=newDeck()
	cardss.printdeck()
}


