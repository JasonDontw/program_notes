package main

import(
	"fmt"
)
	
type bot interface{ //定義一個接口
	getGreeting() string //將名稱為getGreeting()的函式都倒到這個街口
}

type englistBot struct{}
type spanlishBot struct{}


func main(){
	eb := englistBot{}
	sb := spanlishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){ //以接口去執行導入的函數，接口會自行判斷輸入變數需執行哪一個函式
	fmt.Println(b.getGreeting())
}

func (eb englistBot) getGreeting()string{
	return "hi there"
}

func (sb spanlishBot) getGreeting()string{
	return "hola"
}