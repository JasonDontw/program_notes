package main

import(
	"fmt"
)


func main(){
	//var color map[string]string
	//color:=make(map[string]string)
	color := map[string]string{
		"red":"#ff0000",
		"blue":"#4bf745",
	}
	color["black"]="#ffffff" //新增直接撰寫即可不用使用函數
	delete(color, "blue") //刪除元素方法
	fmt.Println(color)
	fmt.Println(color["red"])
	printMap(color)
}


func printMap(m map[string]string){
	for color,hex:=range m{
		fmt.Println("hex code for",color,"is",hex)
	}

}