//Write a program that declares a map of strings to integers and prints out the values in reverse order.

package main

import "fmt"

func main() {
	var maps = map[string]int{}
	var n int
	fmt.Println("Enter the length of the map ")
	fmt.Scanf("%d", &n)
	var num int
	var text string
	fmt.Println("Enter the keys and values: ")
	for i := 1; i <= n; i++ {
		
		fmt.Scanf("%s %d", &text, &num)
	
	maps[text] = num // takes input from user
	}
	fmt.Println(maps)
	//Create a slice to store the keys in reverse order
	keys := make([]string, 0, len(maps))
	for k := range maps {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	//Print out the values in reverse order
	for i := len(keys) - 1; i >= 0; i-- {
		fmt.Printf("%v: %v\n", keys[i], maps[keys[i]])
	}
}
