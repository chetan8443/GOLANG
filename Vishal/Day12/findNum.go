package main
//find elements which are repeating in array
import "fmt"

func main() {
	var arr = []int{10, 30, 30, 20, 10, 20, 50, 10}
	countFreq(arr, len(arr))
}

func countFreq(arr []int, n int) {

	var visited = [8]bool{}
	// Traverse through array elements and
         // count frequencies
	for i := 0; i < n; i++ {
		if visited[i] == true {
			continue
		}
		count := 1
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				visited[j] = true
				count++
			}
		}
		if count != 1 {
			fmt.Println(arr[i])
		}
	}
}
