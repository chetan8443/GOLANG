//Write a program that declares a map of strings to integers and prints out the values in reverse order.

package main

import "fmt"

func main() {
	var maps = map[string]int{}
	var n int
	fmt.Println("Enter the number of values to be added in map.")
	fmt.Scanf("%d", &n)
	fmt.Println("Enter the keys and values: ")
	for i := 0; i < n; i++ {
		var num int
		var text string
		fmt.Scanf("%s %d", &text, &num)
		maps[text] = num // takes input from user
	}

	// changes the key and value
	fmt.Println(maps)
	var rev = map[int]string{}
	for k, v := range maps {
		rev[v] = k
	}
	fmt.Println(rev)
}
