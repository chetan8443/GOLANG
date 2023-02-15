package main
import "fmt"

func sol(num int)int{
	count:=0
	for num>0{
		num /= 10
		count++
	}
	return count
}
 func main(){
	fmt.Println(sol(78))

	fmt.Print("Enter number: ")  
	var input int  
   fmt.Scanln(&input)
   fmt.Println(sol(input))
 }