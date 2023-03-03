package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	
	go OddEven(&wg)
	wg.Wait()

	c := make(chan int, 2)
	go checkOccurence(c) //function uses go routing
	
	fmt.Println("number of occurence of given Char is :",<-c)
	
}

func OddEven(wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	var odd []int
	var even []int
	for _, value := range arr {
		if value%2 == 0 {
			even = append(even, value)
		} else {
			odd = append(odd, value)
		}
	}
	fmt.Println("The Even elements are:", even)
	fmt.Println("The Even elements are:", odd)
}

func checkOccurence(c chan int) {

	fmt.Println("enter word :")
	var word string
	fmt.Scan(&word)
	w := strings.ToLower(word)

	fmt.Println("enter CHAR :")

	var char string
	fmt.Scan(&char)
	cc := strings.ToLower(char)
	var count = 0

	//for loop use for comparing each word from string
	for i := 0; i < len(w); i++ {

		if string(w[i]) == cc {

			count++
		}

	}
	c <- count

}
