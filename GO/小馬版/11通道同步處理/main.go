package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"
	"time"
)

func action(done chan bool) {
	fmt.Println("Loading...")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished.")

	done <- true //向通道發送完成訊息
}

func main() {
	done := make(chan bool)
	go action(done)
	//等待子線程通道消息
	<-done
}