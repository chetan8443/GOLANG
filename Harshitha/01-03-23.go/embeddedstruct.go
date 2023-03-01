package main

import "fmt"

//Usage of embedded structures

type contactinfo struct{
	email string
	postalcode int

}
type person struct{
	firstname string
	lastname string
	contact contactinfo

}

func main(){
	//create a value of type person
	/*Harshitha:=person{lastname:"Harshitha",firstname:"Tangala"}
	Harshitha.firstname="Venkateswarlu"
	fmt.Printf("%+v",Harshitha)
	fmt.Println("\n")
	fmt.Println(Harshitha)*/
	a:=person{
		firstname:"Harshitha",
		lastname:"Tangala",
		contact:contactinfo{
			email:"harshitha@gmail.com",
			postalcode:517325,},
	}
	a.print()
	apointer:=&a
	apointer.update("Indravathi")
	a.print()
}
//updates on a struct
func (pointerToPerson  *person) update(newname string){
	(*pointerToPerson).firstname=newname
}
//The reciever function
func (p person) print(){
	fmt.Printf("%+v",p)
}
