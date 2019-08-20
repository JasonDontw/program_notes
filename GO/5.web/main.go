package main

import(
	"fmt"
	"net/http"
	"os"
	"io"
)
	
type logWriter struct{}

func main(){

	resp, err :=http.Get("http://google.com")
	if err!=nil{
		fmt.Println("error",err)
		os.Exit(1)
	}

	  
	// x:=make([]byte,50000)     //創立一個99999長度的byte陣列
	// resp.Body.Read(x)         //將頁面的原始碼寫進陣列裡
	// fmt.Println(string(x))    //因resp.Body.Read寫入的都是ASCII所以要轉成string

	lw := logWriter{} 
	io.Copy(lw, resp.Body)
}

func(logWriter) Write(bs []byte)(int, error){
	fmt.Println("*********************************")
	fmt.Println(string(bs))
	fmt.Println("*********************************")
	fmt.Println(len(bs))

	return len(bs),nil
}

