package main

import "fmt"

// func main() {

// 	var a int = 14

// 	ptr := &a
// 	fmt.Println(ptr)
// 	fmt.Println(*ptr)

// 	*ptr = 45
// 	fmt.Println(a)
// }

// var pnt *temp

// var x = 123
// var pnt1 *int = &x
// fmt.Println(pnt)
// fmt.Println(*pnt1)
// type temp struct {
// }

// func changevalue(ptr *int) {
// 	*ptr = 67
// }

// func main() {

// 	var x int = 43
// 	ptr := &x
// 	fmt.Println(ptr)
// 	changevalue(ptr)
// 	fmt.Println(x)
// }

// //=========
// func divide(x, y int) (float64, error) {
// 	if y == 0 {
// 		return 0, fmt.Errorf("Cannot divide by zero ")
// 		//return 0, errors.New("Cannot divide by zero")
// 	} else {
// 		return float64(x / y), nil
// 	}

// }

// func main() {

// 	data, err := divide(14, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(data)
// 	}
// }


//doubt 
func main() {

	defer safe()
	const one = 2
	if one != 1 {
		panic("one isn't 1 ")
	}
}
func safe() {
	r := recover()
	if r != nil {
		fmt.Println("Panic is recovred")
	}
}
