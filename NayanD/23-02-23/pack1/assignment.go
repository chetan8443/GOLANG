package pack1

import "fmt"

var Name string = "Nayan Dhoble" // declaring global varaiable

func Swap() { //swap function
	var num1, num2 int
	num1 = 10
	num2 = 20
	fmt.Println("Numbers before swap: \n Num1 =", num1, "\n Num2 =", num2)
	num1, num2 = num2, num1
	fmt.Println("Numbers after swap:\n Num1 =", num1, "\n Num2 =", num2, "\n(Swap using a one-liner syntax)")
}
