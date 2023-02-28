package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Println("ENter string")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	fmt.Println(str)
}
