//Write a program to make a map using two list :

package main

import "fmt"

func main() {

	list1 := []int{
		1, 2, 3,
	}

	list2 := []string{
		"apple", "banana", "cherry",
	}

	fmt.Println(makeMap(list1, list2))

}

func makeMap(l1 []int, l2 []string) map[int]string {
	m := map[int]string{}
	for i, k := range l1 {
		m[k] = l2[i]
	}
	return m
}
