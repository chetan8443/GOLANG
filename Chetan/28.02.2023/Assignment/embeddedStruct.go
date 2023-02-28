package main

import "fmt"

//created student struct
type student struct {
	name     string
	house_no int
	contact  contacts
}

//created contact struct
type contacts struct {
	mobno   int
	zipcode int
}


//created main function
func main() {


 //created obj reference by assigning values to struct variable 
	obj := student{
		house_no: 43,
		name:     "sujay",
		contact: contacts{mobno: 9349934,
			zipcode: 420012},
	}
	// fmt.Print(obj)

	obj.print()
	var objPointer *student= &obj
	objPointer.updateName("chetan")
	obj.print()
}

//created print function
func (p student) print() {
	fmt.Print(p)
}
func (s *student)updateName(newName string){
	s.name=newName
}
