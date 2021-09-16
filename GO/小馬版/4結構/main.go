package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
			 //若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"  //基本套件
)

type user struct { //宣告結構，類似PHP的class
	name string
	password string
	age int
}

type rect struct { //宣告結構，類似PHP的class
	width, height int
}

func (r *rect) area() int { //(r *rect)此語法為指定給結構rect來調用，但是他是使用指標指向此結構來調用
	return r.width * r.height
}

func (r rect) perim() int { //(r rect)此語法為指定給結構rect來調用，但是額外產生一份資料來調用
	return 2 * (r.width + r.height)
}

func main() {
	user1 := user{name: "lucas", password: "12334456", age: 25}
	fmt.Println("結構內容", user1.name, user1.password, user1.age)  

	userP := &user1 //宣告一個指標，指向結構
	fmt.Println("指標內容", userP.name, userP.password, userP.age)  

	userP.name = "dong"
	userP.password = "741850"
	userP.age = 30

	fmt.Println("指標修改後，結構內容", user1.name, user1.password, user1.age)  
	fmt.Println("指標修改後，指標內容", userP.name, userP.password, userP.age)  

	//結構內建函式調用
	r := rect{width: 10, height: 5}
	fmt.Println("普通結構area", r.area())  
	fmt.Println("普通結構perim", r.perim())  

	//指標方法調用
	rp := &r
	fmt.Println("指標結構area", rp.area())  
	fmt.Println("指標結構perim", rp.perim())  
}