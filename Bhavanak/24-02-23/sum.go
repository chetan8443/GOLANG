/*
Given an array of integers, find the sum of its elements.
For example, if the array , , so return .
Function DescriptionComplete the simpleArraySum function in the editor below. It must return the sum of the array elements as an integer.
simpleArraySum has the following parameter(s):
ar: an array of integers
Input Format
The first line contains an integer, , denoting the size of the array.
The second line contains  space-separated integers representing the array's elements.
Constraints
Output Format
Print the sum of the array's elements as a single integer.
Sample Input
6
1 2 3 4 10 11
Sample Output
31
*/
/*package main

import "fmt"

func main() {
	var n, i int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&n)
	var arr [6]int
	for i = 0; i < n; i++ {
		fmt.Println("enter a number")
		fmt.Scanln(&arr[i])
	}
	var sum int = 0
	for i = 0; i < n; i++ {
		sum = sum + arr[i]

	}
	fmt.Println(sum)
}
*/

package main

import "fmt"

func arrsum(arr [6]int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func main() {
	var n, i int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&n)
	var arr [6]int
	for i = 0; i < n; i++ {
		fmt.Println("enter a number")
		fmt.Scanln(&arr[i])
	}
	fmt.Println("The sum of elements", arrsum(arr))
}
