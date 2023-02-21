package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "how are you bro? wassup?"
	fmt.Println(str)
	fmt.Println("Enter charachter for count for that char in above string : ")
	var chr string
	fmt.Scan(&chr)
	fmt.Println(strings.Count(str, chr)) // 4
}
