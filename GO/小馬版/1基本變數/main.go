package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
			 //若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"  //基本套件
)

func main() {
	//字串
	var a = "hello " + "world" //動態定義，變數會直接抓給定值的型態，所以此命名方式一定要給值         
	fmt.Println(a)
	fmt.Println("hello", "world") //印出也可以直接用,隔開
	//整數
	var b,c int = 10, 20 //一次定義多個(數值一定必須定義型態，因為數值型態有多種系統會無法判斷)
	fmt.Println(b, c)
	//布林
	var d = true
	fmt.Println(d)
	//浮點數
	var e float32 //直接定義，若直接定義型態，則可以不用附值
	e = 3.141596
	fmt.Println(e)
	//指標
	var pa *int 
	pa = &b //變量的地址
	fmt.Println(pa) //印出pa指向的地址
	fmt.Println(*pa) //印出pa指向的地址的值
	
	//定義類型+直接賦值
	f := 5 //go項目常見寫法，不需使用var定義
	fmt.Println(f)

	//陣列使用
	var g [5]int
	fmt.Println("g數組", g)

	h := [5]int{1, 2, 3, 4, 5}
	fmt.Println("h數組", h)

	//切片
	var i []string //創建時無定義長度，則定義為切片
	j := make([]string, 3, 5) //切片一般用make創建，元素1為類型，元素2為長度，元素3為容量(預設跟長度一樣)，長度為目前可訪問的個數，容量為還可擴充的數量，超出容量則以容量大小*2擴充
	fmt.Println("i切片", i, len(i), cap(i))
	fmt.Println("j切片", j, len(j), cap(j))
	j = append(j, "a", "b", "c", "d", "e", "f", "g") 
	fmt.Println("j切片擴充", j, len(j), cap(j))

	k := h[3:] //index3之後的數值寫入k中，不含index3
	k = append(k, 'a', 'b', 'c', 'd', 'e', 'f')
	fmt.Println("k切片", k)

	/////////////切片、陣列差別解說/////////////
	//切片不能放不同型態的值，若硬塞系統會自動轉譯成相同型態
	//陣列可以放不同型態
	//切片可以長度可以擴長
	//陣列長度初始就必須固定
	//只要沒有完整定義出長度的都算切片
	//切片不能用 === 之類的判斷是比較
	//陣列能用 === 之類的判斷是比較
	//切片不能與陣列比較
	//陣列在控制長度、內存、格式方面較具優勢，且有編譯安全檢查能提早發現問題
	//陣列特定單個元素訪問的速度較快

	//哈希表、字典(無序的集合)
	l := make(map[string]int) //map[key type]value type
	//var l map[string]int
	fmt.Println("l哈希表", l, len(l))
	l["k1"] = 1
	l["k2"] = 2
	fmt.Println("l哈希表賦值", l, len(l))
	delete(l, "k1")
	fmt.Println("l哈希表刪除", l, len(l))

	m := map[string]int{"k1" : 1, "k2" : 2} //定義+初始化
	fmt.Println("m哈希表", m, len(m))

	//range使用(迭代數組、切片、字典等等用函式)
	n := []int{1, 2, 3}
	sum := 0
	
	for _, num := range n { //_為空白不使用，key, value := 循環變數，使用Range抽出n的key value賦予前值
		sum += num
	}
	fmt.Println("n總和", sum)

	for key := range n { //若只想迭代key，只需一個變數即可
		fmt.Println("n key", key)
	}
	//其餘類型集合用法都一樣

	for _, value := range "abc" { //也可以直接迭代字串
		fmt.Println("字串迭代", string(value))
	}
}