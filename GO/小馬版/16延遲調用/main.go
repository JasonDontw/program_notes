package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import "fmt"

func MyDeferFunc1() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}

func main() {
	MyDeferFunc1()
}

//用於在程式碼過多時，且此動作又必須在最後執行，就能使用defer把動作移至程式碼最前面，方便閱讀
//多defer調用時，會由下往上執行