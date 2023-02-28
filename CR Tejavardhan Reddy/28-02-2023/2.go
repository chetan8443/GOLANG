// Working with maps
package main

import "fmt"

func main() {
	//var colors map[string]string
	colors := map[string]string{
		"red":   "der",
		"White": "etihW",
		"Green": "neerg",
	}
	colors["black"] = "kcalb"
	fmt.Println(colors)
	print(colors)
}
func print(o map[string]string) {
	for color, reverse := range o {
		fmt.Println(color, "=", reverse)
	}
}
