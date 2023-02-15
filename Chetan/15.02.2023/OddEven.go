package main
import "fmt"

func main(){
	
fmt.Print("ENTER A NUMBER :")
var num int;
fmt.Scan(&num)
if num==0{
	fmt.Print("number is 0")
}else if num%2==0{
	fmt.Print("number is EVEN")
}else{
	fmt.Print("number is ODD")
}

}
