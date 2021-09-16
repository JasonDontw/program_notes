package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import "fmt"

func main() {
	fmt.Println("開始")

	defer func ()  {
		if err := recover(); err !=nil {
			fmt.Println("偵測到錯誤")
			fmt.Println("錯誤原因:" , err)
		}	
	}()

	panic("系統崩饋")

	fmt.Println("結束")
}

//panic類似PHP的throw Exception，執行後會強制中斷所有，並一定會執行defer函式
//且於defer函式中，使用recover來捕捉錯誤且處理類似catch