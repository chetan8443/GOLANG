package main
import ("fmt")

func main() {
	fmt.Println("Enter a number for checking : ")
	var num int
	fmt.Scan(&num)									//assign given number at address of num variable
  
	
	if num > 10 {									//checking condition,if it is true then it will enter in block
    		fmt.Println("Num is more than 10.")

			if num > 15 {							//nested if
			fmt.Println("Num is also more than 15.")
			}
			
  }else if(num == 10){
	fmt.Println("Num is 10.")
	}else {											//if niether condition is satisfy then else block will work
    fmt.Println("Num is less than 10.")
  }
}
