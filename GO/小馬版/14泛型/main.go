package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import "fmt"

func f1(val interface{}) {
	fmt.Println(val)
}

func main() {
	var x interface{} = 100
	fmt.Println(x)
	x = "qwe"
	fmt.Println(x)
	x = false
	fmt.Println(x)

	var y interface{} = 100
	z, ok := y.(float32) //此種方法不可隨便轉，如浮點數轉整數式不行的
	fmt.Println(z, ok)

	f1(123)
	f1("asdas")
	f1(true)
}