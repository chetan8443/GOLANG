//Pgm to search for an element in the slice
package main

import "fmt"

func main() {
	var size, num int
	fmt.Print("Enter the size of the slice")
	fmt.Scan(&size)
	var slice1 = make([]int,size)
	fmt.Print("Enter the elements of the slice")
	for i:=0;i<size;i++ {
		fmt.Scan(&slice1[i])
	}
	fmt.Print("Enter element to be searched for ")
	fmt.Scan(&num)
	search(slice1,num)
}
func search(sl [] int,n int) {
	var found byte
	for i :=0;i<len(sl);i++{
		if(sl[i] == n){
			fmt.Println(n,"found at the ",i+1,"th position")
			found = 'y'
			break
		}
		found ='n'
	}
	if found == 'n' {
		fmt.Print(n," is not found")
	}
}