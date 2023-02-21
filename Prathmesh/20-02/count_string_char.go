package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Mary had a little lamb"
	fmt.Println(strings.Count(str, "a")) // 4
}
