package main

import "fmt"

func main() {
	bsort()
}

func bsort() {
	arr1 := []int{15, 43, 3, 24, 1, 6, 7, 8, 2, 4}
	println("array before sorting: ")
	print(arr1)

	for i := 0; i < len(arr1)-1; i++ {
		for j := 0; j < len(arr1)-1; j++ {
			if arr1[j] > arr1[j+1] {
				temp := arr1[j]
				arr1[j] = arr1[j+1]
				arr1[j+1] = temp
			}

		}
	}
	fmt.Println("array after sorting : ")
	print(arr1)

}

func print(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if i == len(arr)-2 {
			fmt.Println(arr[i])
			break
		}
		fmt.Print(arr[i], ",")
	}
}
