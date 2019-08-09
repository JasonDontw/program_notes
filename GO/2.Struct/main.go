package main

import(
	"fmt"
)

type contactInfo struct{
	email string
	zipCode int 
}

type person struct{
	firstname string
	lastname string
	contactInfo  //struct可以繼承struct
}


func main(){
	Jason := person{
		firstname:"Don",
		lastname:"Jason",
		contactInfo:contactInfo{
			email:"ee13751",
			zipCode:429,
		},
	}
	Jason.firstname="Dong"

	
	Jason.updateFirstName("Don")
	Jason.printperson()
	
}

func (p person)updateFirstNameError(newFirstName string){ //當物件使用此函數時是創建一個新的相同物件
	p.firstname = newFirstName 						      //傳進func裡面所以更改不到原本的
}
func (p *person)updateFirstName(newFirstName string){ 
	(*p).firstname = newFirstName 					
}

func (p person)printperson(){
	fmt.Printf("%+v",p) //印出整個結構
}

//補充:slice可以不需用point也能直接從func裡面改到原始值，因為slice本身就是用point指向了