package main

import (
    "fmt"
    "math"
)



func main() {
 
 var a:= testNumber(153)
 if a{
 fmt.Println("this is armstrong number")
 }else{
  fmt.Println("this is Not armstrong number")
 }
 
    
}

func testNumber(n int) bool {
   
    sum := 0
    for i := n; i > 0; i /= 10 {
        digit := i % 10
        sum += int(math.Pow(float64(digit), float64(3)))
    }

  
    return sum == n
}
