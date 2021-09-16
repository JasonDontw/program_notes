package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"
)

func main() {
	message := make(chan string, 3) //定義一個有3個緩衝區的字串通道，當緩衝區消息達到3個時會阻塞當前線程
									//若沒定義緩衝區，則只有1，若消息超出緩衝區數量阻塞，等到其他線程來接收，若沒有其他線程，則會出錯
	message <- "消息1"
	message <- "消息2"
	message <- "消息3"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}