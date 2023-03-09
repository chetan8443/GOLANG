package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	eid   int    `json:"id"`
	ename string `json:"name"`
	age   int    `json:"age"`
}

func main() {
	// employee := []byte(`{
	// 	"eid" : 1001,
	// 	"ename":"Teja",
	// 	"age":20
	// }`)
	// // {eid: 1001, ename: "teja", age: 20},
	// // {eid: 1002, ename: "vardhan", age: 21},
	// // {eid: 1003, ename: "Reddy", age: 22},)
	// // fmt.Printf("%+v\n", employee)
	// // a, _ := json.Marshal(employee)
	// //fmt.Print(a)
	// var emp1 Employee
	// err := json.Unmarshal(employee, &emp1)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Print(emp1)
	var rawJSON = `{
			"id":1001,
			"name":"Teja",
			"age":22
		}`
	var emp Employee
	err := json.Unmarshal([]byte(rawJSON), &emp)
	if err != nil {
		fmt.Print("Error:", err)
	}
	fmt.Printf("%+v\n", emp.ename)
	fmt.Printf("%+v\n", emp.eid)
	fmt.Printf("%+v", emp.age)

}
