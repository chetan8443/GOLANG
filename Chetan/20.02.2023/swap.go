package main
import "fmt"

func main(){
	var num1,num2 int

	fmt.Print("Enter num1 :")
	fmt.Scan(&num1)
	fmt.Print("Enter num2 :")
	fmt.Scan(&num2)
	fmt.Printf("given num1 is %d & num2 is %d \n",num1 ,num2)
	num1=num1+num2
	num2=num1-num2
	num1=num1-num2
	fmt.Printf("After swapping num1 is %d & num2 is %d",num1 ,num2)
	
}
