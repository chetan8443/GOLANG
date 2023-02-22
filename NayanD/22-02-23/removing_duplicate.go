package main

import "fmt"

func main() {
	list := []int{1, 2, 4, 3, 2, 7, 5, 4, 5, 6, 7, 8, 6}

	new_list := []int{}

	for i := 0; i < len(list); i++ {
		count := 1

		for j := 0; j < len(new_list); j++ {
			if list[i] == new_list[j] {
				count++
			}

		}
		if count == 1 {
			new_list = append(new_list, list[i])

		}
	}
	fmt.Println(new_list)
}
