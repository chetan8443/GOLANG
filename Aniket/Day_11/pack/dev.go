package pack

import "fmt"

type Details interface {
	details()
}

type Tax interface {
	taxAmount()
}

func (e employee) details() {
	fmt.Println("Name:", e.name)
	fmt.Println("Designation:", e.designation)
	fmt.Println("Department:", e.department)
	fmt.Println("Is Permanant:", e.isPermanant)
}

func (m manager) details() {
	fmt.Println("Name of Manager:", m.name)
	fmt.Println("Department:", m.department)
}

func (e employee) taxAmount() {
	fmt.Println("Net Amount:", (e.totalAmount - e.taxableAmount))
}

func (m manager) taxAmount() {
	fmt.Println("Net Amount:", (m.totalAmount - m.taxableAmount))
}

func getInfoD(d1 Details, d2 Details) {
	d1.details()
	d2.details()
}

func getInfoT(t1 Tax, t2 Tax) {
	t1.taxAmount()
	t2.taxAmount()
}
