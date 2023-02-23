package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "my new text is this long"
	y := strings.Repeat("#", len(x))
	fmt.Println(y)

}
