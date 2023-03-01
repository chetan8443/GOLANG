//embedded struct example
package main  
import (  
   "fmt"  
)  
type person struct {  
   fname string  
   lname string}  
type employee struct {  
   person  
   empId int  
}  
func (p person) details() {  
   fmt.Println(p, " "+" I am a memer")  
}  
func (e employee) details() {  
   fmt.Println(e, " "+"I am a employee")  
}  
func main() {  
   p1 := person{"Dhiraj", "Lakhane"}  
   p1.details()  
   e1 := employee{person:person{"John", "Ponting"}, empId: 11}  
   e1.details()  
} 