// // Given the arrival time of each student and a threshhold number of attendees, determine if the class is cancelled.
// int k: the threshold number of students
// int a[n]: the arrival times of the n students
//Sample Input

// 2
//testcase 1:
// 4 3
// -1 -3 4 2
//testcase 2:
// 4 2
// 0 -1 2 1
// Sample Output

// YES
// NO
// Explanation:
// For the first test case,k=3 . The professor wants at least 3 students in attendance, but only 2 have arrived on time (-3 and -1) so the class is cancelled.

// For the second test case,k=2 . The professor wants at least 2 students in attendance, and there are 2 who arrived on time (0 and -1). The class is not cancelled.
package main

import "fmt"

func main() {
	fmt.Println("Enter minimum number of students:")
	var k int
	fmt.Scanln(&k)
	fmt.Println("Enter total number of students:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Println("Enter the time delay/time early by minutes of each student")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	var ontime = 0
	for _, num := range arr {
		if num <= 0 {
			ontime++
		}
	}
	if ontime >= k {
		fmt.Println("The class will continue")
	} else {
		fmt.Println("The class will be suspended")
	}
}
