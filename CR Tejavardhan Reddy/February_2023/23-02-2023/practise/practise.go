package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var str string = "Tejavardhan"
	char := strings.Split(str, "")
	fmt.Println(char)
	word := strings.Join(char, "")
	fmt.Println(word)
	randNumber()
}

func randNumber() {
	var Range int
	fmt.Println("Enter the range:")
	fmt.Scanln(&Range)
	for i := 0; i < Range; i++ {
		fmt.Println(rand.Intn(Range))
	}
}
