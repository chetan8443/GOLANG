package main

import "fmt"

const i int = 3
const Pi float64 = 3.14
const s bool = true
const n string = "yes"

func main() {
	const World = "guys"
	fmt.Println("Nithin:Hey", World)
	fmt.Println("do you know that we can take only take ", Pi, "days of leaves in the span of", i, " mnths")
	fmt.Println("Navneeth:Is it true")
	fmt.Println("Nithin:", s)
	fmt.Println("Nithin:so we can plan to go out those days")
	fmt.Println("Navneeth:", n)

}
