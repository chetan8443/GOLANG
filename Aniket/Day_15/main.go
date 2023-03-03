package main

import "fmt"

func main() {
	var arr = []int32{-1, -2, 0, 1, 2}
	plusMinus(arr)

}
func plusMinus(arr []int32) {
	// Write your code here
	var cpos int = 0
	var cneg int = 0
	var czero int = 0
	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] > 0 {
			cpos = cpos + 1
		} else if arr[i] == 0 {
			czero = czero + 1
		} else {
			cneg = cneg + 1
		}
	}
	fmt.Printf("%.8f", float64(cpos)/float64(len(arr)))
	fmt.Println(cneg / len(arr))
	fmt.Println(czero / len(arr))
}
