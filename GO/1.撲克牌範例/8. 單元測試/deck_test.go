//用於使用go test 測試deck.go
//創建一個新檔在後面加上_test，然後開始寫要測試檔案的腳本

package main

import ("testing"
		"os"
)
func Test_newDeck(t *testing.T){ //在要測試的函數名稱前加上Test
	d := newDeck()

	if len(d)!=16{
		t.Errorf("error len %d",len(d))
	}

	if d[0]!="Ace of Spades" {
		t.Errorf("error first card %v",d[0])
	}

	if d[0]!="Ace of Spades" {
		t.Errorf("error first card %v",d[0])
	}
}

func Test_saveToFileAndnewDeckFromFile(t *testing.T){
	os.Remove("_decktesting") //將此檔案移除

	d:=newDeck()  //創建一個新牌組
	d.saveToFile("_decktesting") //存成叫"_decktesting"的檔案

	loadedDeck := newDeckFromFile("_decktesting") //以檔案方式載入牌組

	if len(loadedDeck) != 16 { //驗證載入牌組是否正確
		t.Errorf("error len %d",len(d))
	}
	os.Remove("_decktesting") //測試完後也要將此檔案移除

}


