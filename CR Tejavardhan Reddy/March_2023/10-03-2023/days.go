// Golang code to find the number of good days in between range
// Explanation:
// days 20, 21, 22, and 23.
// We perform the following calculations to determine which days are beautiful:
// -->Day 20 is beautiful because the following evaluates to a whole number:
// |20-02|/6=3
// -->Day 21 is not beautiful because the following doesn't evaluate to a whole number:
// |21-12|/6 = 1.5
// -->Day 22 is beautiful because the following evaluates to a whole number:
// |22-22|/6=0
// -->Day 23 is not beautiful because the following doesn't evaluate to a whole number:
// |23-32|/6 = 1.5
// Only two days, 20 and 22, in this interval are beautiful. Thus, we print 2 as our answer.

package main

import "fmt"

func main() {
	fmt.Println("Enter the starting day:")
	var start int
	fmt.Scanln(&start)
	fmt.Println("Enter the ending day:")
	var end int
	fmt.Scanln(&end)
	fmt.Println("Enter the k value:")
	var k int
	fmt.Scanln(&k)
	count := 0
	for i := start; i <= end; i++ {
		a := reverse(i)
		diff := i - a
		if diff < 0 {
			diff *= -1
		}
		if diff%k == 0 {
			count++
		} else {
			continue
		}
	}
	fmt.Println("The number of good days:", count)
}

func reverse(i int) int {
	var rem, rev int
	for i > 0 {
		rem = i % 10
		rev = rev*10 + rem
		i = i / 10
	}
	return rev

}

// func goodDay(i,a,k int) {
// 	diff:=i-a
// 	if diff<0 {
// 		diff*=-1
// 	}
// 	if diff%k==0 {
// 		return "The number of good days are:"
// 	}
// }
