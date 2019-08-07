package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
			 //若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt" //基本套件
)


func main() { //程式開始執行的入口
	var card string //若無給定數值且又要規定型別的定義方式
	var s = 1       //動態定義，變數會直接抓給定值的型態，所以此命名方式一定要給值           
	var a,b int     //一次定義多個
	x := "Ace of spades"  //此種定義方法會直接抓取給定的值的型態且無法變更型態，等同於 var x string
	cards:=newCard()
	
	c := []string{newCard(),"Diamond","good"} //c為一個陣列且是字串的陣列，若無規定大小則可擴充
	c = append(c,"add") //在陣列C後面新增
	d := [4]int{}  //若有規定大小則不可擴充
	
	fmt.Println(card)
	fmt.Println()/////
	fmt.Println(s+a+b)
	fmt.Println()/////
	fmt.Println(x)
	fmt.Println()/////
	fmt.Println(cards)
	fmt.Println()/////
	fmt.Println(d[0])
	fmt.Println()/////
	for i,cc:=range c{     //此種for第一個為索引值第二個為C陣列依序的值
		fmt.Println(i,cc)
	}


}

func newCard()string{ //後面的string定義了return的型別
	return "five of diamonds"
}
