package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"
	"time"
)

func player(p_name string, p_chan chan int) {
	for {
		val, ok := <-p_chan
		if ok {
			fmt.Println(p_name, val, ok)
		} else {
			break
		}
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	defer close(chan1) //結束後要關閉通道

	go player("玩家一", chan1)
	go player("玩家二", chan2)

	i := 1
	for i <=5 {
		chan1 <- i
		chan1 <- i * 100
		i++
		time.Sleep(time.Second)
	}
}
//協程之間盡量通過主線程來溝通，避免協程之間互相溝通，設計不好會導致不必要的錯誤