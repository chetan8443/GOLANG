package main

import "fmt"

func main() {
   var b int = 15
   var a int


   /* for loop execution */
   for a := 0; a < 10; a++ {
      fmt.Printf("value of a: %d\n", a)
   }
   fmt.Println("=====2nd For Loop=======")
   for a < b {
      a++
	 
      fmt.Printf("value of a: %d\n", a)
   }
  
}
