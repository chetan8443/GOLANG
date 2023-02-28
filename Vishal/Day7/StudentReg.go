// created student registartion using struct of slice
package main

import (
	"fmt"

)

func main() {
	sliceOfstruct()
}

func sliceOfstruct() {
	type student struct {  //created struct
		name   string
		rollNo int
		class  string
	}

	


	fmt.Printf("Enter Number of student You wants to register: ")
	var n int
	fmt.Scanln(&n)
	students := []student{{   // declared slice of struct
		name:"",
		class:"",
		rollNo:0,},
	}
     
	for i := 0; i <= n-1; i++ {
		var nam string
		var className string
		fmt.Printf("Enter Name of student :")
		fmt.Scanln(&nam)
		var stud student
		stud.name=nam
		stud.class=className
		stud.rollNo=i+1
		students=append(students, stud)    //adding values taken from console to slice

	}

	fmt.Println("Here s list of registerd students")   // Printing all the records from the struct
	for _, v := range students {
		fmt.Printf("Name:%s \nrollNo:%d \nClass:%s\n",v.name,v.rollNo,v.class)
		fmt.Println("---------------------")
	}
}
