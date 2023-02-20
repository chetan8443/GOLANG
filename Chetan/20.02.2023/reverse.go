package main
import "fmt"

func main(){
	var num int
	fmt.Println("Enter number to be reverse : ")
	fmt.Scan(&num)
	revers(num)
}
func revers(n int){

	for n!=0{
		rem:=n%10
		n=n/10
		fmt.Print(rem)
	}


}
