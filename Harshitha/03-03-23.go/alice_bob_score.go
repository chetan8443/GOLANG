package main

import "fmt"

func main() {
	var a = [...]int32{1, 2, 3}
	var b = [...]int32{2, 2, 2}
	var n int = len(a)
	var alice, bob int32 = 0, 0
	var i int = 0
	for i = 0; i < n; i++ {
		if a[i] > b[i] {
			alice = alice + 1
		} else if a[i] == b[i] {
			continue
		} else {
			bob = bob + 1
		}
	}
	var result []int32
	result = append(result, alice, bob)
	fmt.Println(result)
}
