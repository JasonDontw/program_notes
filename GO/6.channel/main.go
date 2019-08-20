package main

import(
	"fmt"
	"net/http"
)
	

/////////////////正常寫法///////////////////////
/*
func main(){
	links:=[]string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}
	
	for _,link:=range links{
		checkLink(link)

	}

}

func checkLink(link string){
	_,err:=http.Get(link)
	if err != nil{
		fmt.Println(link,"might be down!")
		return
	}
	fmt.Println(link,"is up")
}
*/

///////////////////////平行處理寫法////////////////////////////

func main(){
	links:=[]string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}
	
	c:=make(chan string) //創建一個Channel

	for _,link:=range links{
		go checkLink(link,c) //加上GO會讓函式不需等上個函式結束就開另一個Routine平行處理
	}

	for i:=0;i<len(links);i++{
		fmt.Println(<-c)
	}
}

func checkLink(link string,c chan string){
	_,err:=http.Get(link)
	if err != nil{
		fmt.Println(link,"might be down!")
		c <-"I think"
		return
	}
	fmt.Println(link,"is up")
	c <- "its up"
}