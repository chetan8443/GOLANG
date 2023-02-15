package main
import "fmt"

func sol(num int)int{
	result:=0
	for num>0{
		result = result + num % 10
		num = num / 10
	}
	return result
}

func main(){
	fmt.Println(sol(78))

	//one with input from cmd

	fmt.Print("Enter number: ")  
	var input int  
   fmt.Scanln(&input)

   fmt.Println(sol(input))
}