// given an array of integers and you need to print te value associated with the even index
package main

import "fmt"

func main() {
	fmt.Print("Enter the size:")
	var size int
	fmt.Scanln(&size)
	fmt.Println("Enter the elements:")
	var arr = make([]int, size)
	//var res = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	for i, ele := range arr {
		if i < 2 {
			continue
		} else {
			a := prime(i, size)
			//fmt.Println(a)
			if a {
				fmt.Println("The numbers that are in prime index are:", ele)
				//res = append(res, ele)
			} else {
				continue
			}
		}
	}
	//fmt.Println(res)

}
func prime(i, size int) bool {
	flag := false
	for k := 2; k <= size/2; k++ {
		//fmt.Println(k)
		if i%k == 0 {
			flag = true
			break
		} else {
			flag = false
			break
		}
	}
	if flag {
		return true
	} else {
		return false
	}
}
