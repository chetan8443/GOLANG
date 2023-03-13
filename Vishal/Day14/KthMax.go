package main
// find kth max element from array
import (
	"fmt"
	"sort"
)

func main() {
	test()
}

func test() {
	var arr = []int{11, 13, 2, 6, 7, 4, 5}

	var input int
	fmt.Printf("enter Kth element to find : ")
	fmt.Scanln(&input)
	sort.Ints(arr)  // sorting array
	
	ab:=arr[input-1]
	fmt.Printf("%d max element is %d",input,ab) // printing output
}
