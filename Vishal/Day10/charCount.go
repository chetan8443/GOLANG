package main
// Program to count number of Occurrences in string
import (
	"fmt"
	
	"sync"
)

func main() {
	fmt.Println("execution of main started")

	var wg sync.WaitGroup   // waitgroup declaration

	wg.Add(1)
	charCounter(&wg)
	wg.Wait()
	fmt.Println("execution of main Finished")
}
func charCounter(wg *sync.WaitGroup)  {
	var res =map[string]int {}  // declaring map to store result 
	var ab="testest"   
      
	for i := 0; i < len(ab); i++ {  // for loop to iterate over a string
		var count = 0
		var a string= string( ab[i])
		for j:=0 ;j<len(ab); j++{
		 
		   if a==string( ab[j]){  
			 count++
		   }
		}
		res[a]=count
	}

	fmt.Println(res)   //printing result
	wg.Done()
}
