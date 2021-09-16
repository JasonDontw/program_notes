package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"
	"os"      //用於執行CMD指令
	"os/exec" //用於執行CMD指令
	"time"
)

func sayHelo(name string) {
	for i := 1; i <=5; i++ {
		fmt.Println("Helo", name, ":", i)
	}
}

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// //同步執行
	sayHelo("koma")
	// //異步執行
	go sayHelo("lucas")
	go sayHelo("dong")
	go sayHelo("xin")
	go sayHelo("yin")

	//匿名函數，異步執行
	go func(msg string) {
		fmt.Println("this is a", msg)
	}("lesson")
	
	time.Sleep(time.Second) //主線程要加上等待，不然主線程的執行是不等待子線程的，若主線程節，則所有子線程都會被暴力終結
}