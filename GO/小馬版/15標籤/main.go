package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import "fmt"

func main() {
	main: //類似goto用法
	for {
		for {
			fmt.Println("開始")
			break main //demo時請使用三種情境比較，空、break、break main
		}
		fmt.Println("循環中")
	}
}