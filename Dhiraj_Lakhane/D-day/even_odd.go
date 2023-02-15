package main
import "fmt"
func main() {  
	fmt.Print("Enter number: ")  
	var input int  
   fmt.Scanln(&input)  
	fmt.Print(input)  
	/* check the boolean condition */  
	if( input % 2==0 ) {  
	   /* if condition is true then print the following */  
	   fmt.Printf(" is even\n" );  
	} else {  
	   /* if condition is false then print the following */  
	   fmt.Printf(" is odd\n" );  
	}  
 }  