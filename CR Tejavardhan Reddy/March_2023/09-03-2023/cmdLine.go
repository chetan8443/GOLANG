package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Println(os.Args)
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Println(arg)
		byteCode := []byte(arg)
		fmt.Println(byteCode)
		for i, ch := range byteCode {
			if ch >= 65 && ch <= 90 {
				log.Print("The Uppercase letter:", string(ch), " at index ", i)
			} else if ch >= 96 && ch <= 122 {
				log.Print("The Lowecase letter:", string(ch), " at index ", i)
			} else {
				log.Print("The Special charater:", string(ch), " at index", i)
			}
		}
	}
}
