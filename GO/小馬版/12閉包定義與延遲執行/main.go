package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import "fmt"

func MyLaterFunc() func(string) string {
	var _oldVal string

	return func (newVal string) string {
		result := _oldVal
		_oldVal = newVal

		return result
	}
}

func main() {
	f1 := MyLaterFunc()
	result := f1("helo")
	fmt.Println(result)
	result = f1("go")
	fmt.Println(result)
}