package main

import "fmt"

/*func main() {
	mem := member()
	var arr = [10]int{}
	// var i int
	fmt.Println("Enter the elements into the array")
	for i := 0; i < len(arr); i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	var d map[int]int
	for i := 0; i < len(arr); i++ {
		if mem(arr[i]) {
			return arr[i]
		} else {
			d[arr[i]] = 1
		}
	}
}
func member(arr [i]int) func() bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == d[i] {
			return func() bool {
				return true
			}
		}

	}
	return func() bool {
		return false
	}
}*/
//It can be done by the usage of two for loops easily
func main() {
	var arr [6]int
	var item int = 0
	var index int = 0

	fmt.Printf("Enter array elements: \n")
	for i := 0; i <= 5; i++ {
		fmt.Scanln(&arr[i])
	}

	index = -1
	for i := 0; i < 5; i++ {
		for j := i + 1; j <= 5; j++ {
			if arr[i] == arr[j] {
				item = arr[j]
				index = j
				goto OUT
			}
		}
	}

OUT:
	fmt.Println("The element which is repeating for the first time is:")
	if index != -1 {
		fmt.Println(item, index)
	} else {
		fmt.Print("There is no repeated element")
	}
}
