package main

// /1.Program to create an embedded struct
import "fmt"

type departmentinfo struct {
	deptid   int
	deptName string
}

type employee struct {
	eid   int
	ename string
	dept  departmentinfo
}

func main() {
	e1 := employee{
		eid:   101,
		ename: "Teja",
		dept: departmentinfo{
			deptid:   12,
			deptName: "Golang",
		},
	}
	fmt.Printf("%+v", e1)
}
