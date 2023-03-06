package main
import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var n int
	fmt.Scan(&n)

	go EvenNumbers(n, &wg)
	go OddNumbers(n, &wg)

	wg.Wait()
}
func EvenNumbers(n int, wg *sync.WaitGroup) {
	fmt.Println("")
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			fmt.Print(i," ")
		}
	}
	defer wg.Done()
}
func OddNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
	}
}

