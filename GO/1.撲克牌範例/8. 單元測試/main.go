package main 

import (

//	"io/ioutil"

)


func main() {
	cards := newDeckFromFile("my_cards")//正常此函數要放檔案路徑
										//因為在同一資料夾所以直接用名稱即可
	cards.shuffle()
	cards.printdeck()

}
