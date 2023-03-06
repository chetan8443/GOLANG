// 2.Program to create a slice of and map of structs
package main

import "fmt"

type employee struct {
	eid    int
	ename  string
	deptId string
	empSal int
}

func main() {
	e1 := employee{eid: 1001, ename: "Teja", deptId: "1AC", empSal: 20000}
	e2 := employee{eid: 1002, ename: "Vardhan", deptId: "1AB", empSal: 25000}
	e3 := employee{eid: 1003, ename: "Reddy", deptId: "2BC", empSal: 20000}
	// fmt.Printf("%+v\n%+v\n%+v\n%T", e1, e2, e3, e1)
	// fmt.Println(e1.deptId)
	// for k, v := range e1 {
	// 	fmt.Println(k, "=", v)
	// }
	var slice = make([]employee, 0)
	//fmt.Println(slice)
	slice = append(slice, e1, e2, e3)
	for _, j := range slice {
		fmt.Println(j)
	}
	dict := map[int]employee{1: e1, 2: e2, 3: e3}
	for k, v := range dict {
		fmt.Printf("%v=%+v\n", k, v)
	}
	//fmt.Print(dict)
}
