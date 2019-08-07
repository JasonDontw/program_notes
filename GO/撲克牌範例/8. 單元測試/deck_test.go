//用於使用go test 測試deck.go
//創建一個新檔在後面加上_test，然後開始寫要測試檔案的腳本

package main

import ("testing"
	
		)
func Test_newDeck(t *testing.T){ //在要測試的函數名稱前加上Test
	d := newDeck()

	if len(d)!=16{
		t.Errorf("error %d",len(d))
	}

	if d[0]!="Ace of Spades" {
		t.Errorf("error first card %v",d[0])
	}

	if d[0]!="Ace of Spades" {
		t.Errorf("error first card %v",d[0])
	}
}


