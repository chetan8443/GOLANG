package main
  
import "fmt"
  
// taking a structure
type Employee struct {
  
    // taking variables
    name  string
    empid int
}
  
// Main Function
func main() {
  
    // creating the instance of the
    // Employee struct type
    emp := Employee{"ABC", 19078}
    emp1 := Employee{"Harshitha", 100000}
    
  
    // Here, it is the pointer to the struct
    pts := &emp
    pts1:=&emp1
  
    // displaying the values
    fmt.Println(*pts)
    fmt.Println(pts1)
    // updating the value of name
    pts.name = "XYZ"
    fmt.Println(pts.empid)
  
    fmt.Println(*pts)
  
}
