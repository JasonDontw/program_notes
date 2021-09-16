package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
			 //若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"  //基本套件
)

/////////初始化函式//////// 
func init() { //初始化函式是GO內建的函式，定義完後，GO就會在程式開始時自動執行，且可以定義多個
	//do something
}
func init() {
	//do something
}
/////////普通函式//////////
func plus(a int, b int) int { //分開定義類型，最後一格類型為返回類型
	return a + b
} 

func plus2(a, b, c int) int { //統一定義類型
	return a + b
} 

func calABCD(a, b int) (int, int, int, int) { //返回值不可統一定義
	return a+b, a-b, a*b, a/b
} 

func sum(nums ...int) int { //傳入一個叫nums的集合，不限各數所以用...表示
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func sumByRecursion(num int) int { //遞迴求和
	if num == 1 {
		return num
	}

	return sumByRecursion(num-1) + num
}
/////////函式返回函式//////////
func returnDatabaseType() func(string) string {
	return func(name string) string{
		return "helo" + name
	}
}

/////////回調函式//////////
func updateDB(callback func()) {
	fmt.Println("數據處理中....")
	///////處理邏輯/////////
	fmt.Println("數據更新完成....")
	callback()
}

func main() {
	/////////普通函式//////////
	result := plus(10, 20)
	fmt.Println("plus 10+20=", result)

	result = plus2(10, 20, 30)
	fmt.Println("plus2 10+20+30=", result)  

	a, b, c, d := calABCD(20, 10)
	fmt.Println("calABCD 四則運算", a, b, c, d)  

	result = sum(10, 20, 30, 40)
	fmt.Println("sum 不限個數相加", result)  
	
	result = sumByRecursion(10)
	fmt.Println("sumByRecursion 10遞迴求和", result)  

	//////////匿名函式/////////
	//變數方法
	f1 := func (x, y int) int {
		return x + y
	}

	sum := f1(10, 20)
	fmt.Println("匿名函式變數法", sum)

	//直接執行方法
	sum = func (x, y int) int {
		return x + y
	}(10, 20)

	fmt.Println("匿名函式直接執行", sum)

	/////////函式返回函式/////////
	resultFunc := returnDatabaseType()
	x := resultFunc("lucas")
	fmt.Println("函數返回函數", x)  

	//////////回調函式///////////
	fmt.Println("主線程開始...")

	updateDB(func() { //此函式正常來說必須使用異步處理，但還未說明到併發及通道，所以先用模擬的
		fmt.Println("回調函式")
	})

	fmt.Println("主線程結束")
}