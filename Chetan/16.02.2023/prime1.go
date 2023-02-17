package main
import "fmt"
func main(){

	fmt.Print("enter a num :")
	var num int
	var m,flag int=0,0
	fmt.Scan(&num)
	m=num/2
	for i:=2;i<=m;i++{
		if(num%i==0){
			fmt.Print("Number is not prime")    
			flag=1;    
			break;    
		}
	}
	if(flag==0) {
		fmt.Print("Number is prime")  
	}   
	  
  
 }    
