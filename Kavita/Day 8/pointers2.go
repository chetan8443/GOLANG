package main

import "fmt"

type company struct {
	name string
}

func main() {
	p := []string{"IBM", "Infobell", "TCS"}

	//fmt.Println(*&p)

	ptr := &p[0]
	*ptr = "Infosys"
	fmt.Println(p)
}
