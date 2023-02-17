package main
import "fmt"

func isPrime(num int)bool{
	var i int
	var b bool=true
	if(num<2){
		b=false
	}else{
		for i=2;i<num;i++{
			if(num % i==0){
				b=false
				break
			}
		}
	}
	return b
}
func main(){
	// var sol bool=isPrime(11) ? "Prime number":"Not Prime number"
	fmt.Print("Enter number: ")  
	var input int  
   fmt.Scanln(&input)
	if(isPrime(input)){
		fmt.Println("given number is prime number")
	}else{
		fmt.Println("not prime number")
	}
}