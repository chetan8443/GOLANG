package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Packages are started")
	str1 := "Teja"
	str2 := "Reddy"
	num := 4.0
	//a := []byte(str)
	fmt.Println("strings package", strings.Compare(str1, str2))
	fmt.Println("mathpackages", math.Sqrt(num))
	reader := bufio.NewReader(os.Stdin)
	res, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	ip := "my name is teja"
	b, _ := json.Marshal(ip)
	fmt.Println(b)
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(response)
}
