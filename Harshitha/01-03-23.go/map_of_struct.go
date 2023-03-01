package main

import "fmt"
// Golang program to show how to
// use structs as map keys
  
//declaring a struct
type Address struct {
    firstname  string
	lastname string
	sex string
	phonenumber int
    city    string
    pincode int
}
  
func main() {
  
    // Creating struct instances
    a2 := Address{firstname:"Teja",lastname:"C R",sex:"Male",phonenumber:9087654,city:"Hindupur",pincode:513423}
    a1 := Address{firstname:"bhavana",lastname:"Koppula",sex:"Female",phonenumber:90876589,city:"Nandyala",pincode:519087}
    a3 := Address{firstname:"Faiza",lastname:"Bukka",sex:"Female",phonenumber:5463728,city:"Madanapalee",pincode:515432}
  
    // Declaring a map
    // Declaring and initialising
    // using map literals
    sample := map[int]Address{1:a1,2:a2,3:a3}
	//since we cannot apply slices on maps
    fmt.Println(sample)
}