package main
import "fmt"



func main(){
fmt.Println("Enter 2 numbers :")
var num1 int;
var num2 int;

fmt.Scan(&num1,&num2)
if(num1>num2){
	fmt.Print(num1," Number is greater than ",num2)
}else{
fmt.Print(num2," Number is greater than ",num1)
}


}
