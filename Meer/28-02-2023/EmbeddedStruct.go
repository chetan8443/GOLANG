// Program to create an embedded struct
package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 8,
		},
		str: "Meer",
	}

	fmt.Printf("co = %+v\n", co)

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

}
