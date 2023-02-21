package main

import "fmt"

const a bool = true       //declaration & intitialization of bool variable
const b int = 47          //declaration & intitialization of integer variable
const c string = "nanna"  //declaration & intitialization of string variable
const d float64 = 12.2356 //declaration & intitialization of float variable

func main() {

	//printing all the constants ,After initializing constatnt variable we can't reinitialize them
	fmt.Println("Boolean:", a, "\n", "integer:", b, "\n", "String :", c, "\n", "Float64:", d, "\n")

}
