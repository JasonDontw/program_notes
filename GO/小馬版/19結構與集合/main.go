package main //package像是一個群組，package main是指讓此檔案屬於main這個群駔
//若要執行.go一定要從package main執行，其他package不能單獨執行只能掛在main底下才能執行

import (
	"fmt"
)

type Book struct {
	name string
	author string
	price int
}

type BookStore struct {
	booklist []*Book
}
func main() {
	book1 := Book{name: "vue學習", author: "a", price: 20}
	book2 := Book{name: "go學習", author: "b", price: 30}

	bookStore := BookStore{}
	bookStore.booklist = append(bookStore.booklist, &book1, &book2)

	for i, book := range bookStore.booklist {
		fmt.Println(i, book)
	}
}

//指標補充
//*[]book是一個指標，指向book陣列
//[]*book是一個陣列裡面放指向book的指標