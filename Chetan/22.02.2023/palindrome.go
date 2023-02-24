package main
import "fmt"

var sum,temp int =0,0

func main(){
	var num int
	fmt.Println("Enter number to be reverse : ")
	fmt.Scan(&num)
	temp=num
	checkPalindrome(num)
	
}
func checkPalindrome(n int){
	
	
	for n!=0{               // Checking codition for loop ,if it's true then only loop will work

		rem:=n%10
		sum=(sum*10)+rem
		
		n=n/10
		
	}
	
	if(temp==sum){						//checking codition of given number,is same with reverse number
	fmt.Println("palindrome number")
}else
{
	fmt.Println("not palindrome")
}
	



}
