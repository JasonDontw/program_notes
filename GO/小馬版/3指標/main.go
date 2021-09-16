package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
			 //若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"  //基本套件
)

func byValue(ivalue int) {
	ivalue = 0
}

func byPointer(iptr *int) {
	*iptr = 0
}

func main() {
	i := 10
	fmt.Println("initial", i)  

	byValue(i)
	fmt.Println("byValue", i)

	byPointer(&i)
	fmt.Println("byPointer", i)
	fmt.Println("pointerAddress", &i)

	//普通修改及指標修改差別
	//普通修改傳入參數後是會創建另一個變數去接，所以不會動到原始值
	//指標修改傳入的參數是數值得位址，修改時會直接去改位址的值，所以會變更到原始的值



	/////指標詳細說明//////
	x := [5]int{1, 2, 3, 4, 5}
	var y *int = &x[0] //單一指標指向陣列第一個，補充:若直接y := &x[0]賦予位址，y會被視為指標類型
	var z *[5]int = &x //陣列指標指向整個陣列

	fmt.Println("x[0]陣列位址", y)  
	fmt.Println("x[0]陣列的值", *y)
	fmt.Println("y指標本身位址", &y)
	fmt.Println("x陣列所有位址", z)  //待研究
	fmt.Println("x陣列所有值", *z)  
	fmt.Println("z指標本身位址", &z)  
	fmt.Println("x[0]陣列位址", &z[0])  
}