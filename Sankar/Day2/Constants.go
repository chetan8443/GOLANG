package main

import "fmt"

//constant declaration 
const phone = "One Plus"
const model_num int = 21345


func main() {
	fmt.Println("printing constant")
	fmt.Println(phone)
	fmt.Println(model_num)
	model_num := 54322
	fmt.Println("new model or changed mode number : " , model_num)
}
