package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
			 //若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"  //基本套件
	"math"
	"reflect"
)

type geometry interface { //定義幾何圖形接口
	area() float64
	perim() float64
}

type rect struct { //定義矩形結構
	width, height float64
}

type circle struct { //定義圓形結構
	radius float64
}

//矩形結構函式
func (r rect) area() float64 { //若要使用接口，函式名稱必須與接口相同
	return r.width * r.height
}

func (r rect) perim() float64 { //若要使用接口，函式名稱必須與接口相同
	return 2 * (r.width + r.height)
}

//圓形結構函式
func (c circle) area() float64 { //若要使用接口，函式名稱必須與接口相同
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 { //若要使用接口，函式名稱必須與接口相同
	return math.Pi * c.radius
}

//計算函式，參數為幾何圖形接口
func measure(g geometry) { //(g geometry)用接口去接收結構，接口會去分析結構自動調用相對應的函數
	fmt.Println(reflect.TypeOf(g), g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 4, height: 5}
	c := circle{radius: 10}

	//計算面積及周長
	measure(r) 
	fmt.Println()
	measure(c)
}