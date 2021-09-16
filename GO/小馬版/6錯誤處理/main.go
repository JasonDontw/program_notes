package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
			 //若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"  //基本套件
	"errors"
)

//使用內建錯誤套件寫法
func f1(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("參數錯誤")
	}

	return 2 * num, nil
}

//使用結構自訂義錯誤寫法
type error interface { //此接口已有內建可以不寫
    Error() string
}

type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string { //"Error"此函式是內定好的，若return類型為error結構則會自動跑此函式
	return fmt.Sprintf("%d -> %s", e.arg, e.msg)
}

func f2(num int) (int, error) {
	if num < 0 {
		return -1, &argError{arg: num, msg: "參數不能為負值"}
	}

	return 2 * num, nil
}
//f2另一種寫法
// func f2(num int) (int, string) {
// 	if num < 0 {
// 		err := argError{arg: num, msg: "參數不能為負值"}
// 		return -1, err.Error()
// 	}

// 	return 2 * num, ""
// }

func main() {
	//f1
	result, err := f1(10)
	fmt.Println("f1正數", result, err)
	result, err = f1(-10)
	fmt.Println("f1負數", result, err)

	//f2
	result, err = f2(10)
	fmt.Println("f2正數", result, err)
	result, err = f2(-10)
	fmt.Println("f2負數", result, err)
}