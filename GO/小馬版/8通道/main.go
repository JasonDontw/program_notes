package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	message := make(chan string) //定義一個字串通道，通道本身就是指標類型

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Second)
			message <- (strconv.Itoa(i) + ".helo channel.")
		}
	}()

	//接收通道消息
	result := ""
	result = <-message //取得通道資料時，會等待到通道有值之後才會繼續往下
	fmt.Println(result)
	result = <-message
	fmt.Println(result)
	result = <-message
	fmt.Println(result)
}