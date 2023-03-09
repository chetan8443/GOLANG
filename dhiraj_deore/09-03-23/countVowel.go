package main

import "fmt"

func main() {
	fmt.Println(vowels("count vowels"))
}

func vowels(s string) int {
	vow := []string{"a", "o", "i", "u", "e"}
	count := 0
	for _, v := range s {
		for _, v1 := range vow {
			if string(v) == v1 {
				count++
				break
			}
		}
	}
	return count
}
