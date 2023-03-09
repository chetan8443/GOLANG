package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the sentence")
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	//fmt.Print(str)
	marshal(str)
	//log.Print(str)
}
func marshal(str string) {
	res, err := json.Marshal(str)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(res)
	fmt.Print(string(res))
}
