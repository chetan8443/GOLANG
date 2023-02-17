package main
import "fmt"

func main(){
	fmt.Print("Enter a number  to get factorial : ")
	var fact int=1
	var num int
	fmt.Scan(&num)

	for i:=1;i<num;i++{
		
		fact=fact*(i+1)
		
	}
	fmt.Printf("Factorial of %d is: %d",num,fact)
	
	


}
