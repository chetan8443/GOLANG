package main
// program to print even and  oddd numbers using waitgroup
import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Execution of main started")

	var wg sync.WaitGroup  // creating a wait group 
	wg.Add(1)
  go numbersPrinting(&wg)  // calling funcion with address of wait group
  wg.Wait()

  
  fmt.Println("Execution of main Ended")
}

func numbersPrinting(wg *sync.WaitGroup)  {
	var odd=[] int{}
	fmt.Println("Even numbers are :")
	for i := 0; i < 20; i++ {
		if i % 2 == 0{    // checking even numbers
		fmt.Printf("%d,",i)
	}else {
		odd=append(odd, i)  // storing odd element in an array
	}
	
	}
	fmt.Println("\n Odd numbers are :")
	fmt.Println(odd)
	wg.Done()
}

