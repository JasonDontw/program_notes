package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"
	"strconv"
)

func main() {
	message := make(chan string) //定義一個字串通道

	go func() {
		for i := 1; i <= 5; i++ {
			if i == 5 {
				message <- ""
			} else {
				message <- (strconv.Itoa(i) + ".helo channel.")
			}
		}
	}()

	//接收通道消息
	for result := range message{ //無限次數接收
		if result == "" {
			break
		} else {
			fmt.Println(result)
		}
	}
}