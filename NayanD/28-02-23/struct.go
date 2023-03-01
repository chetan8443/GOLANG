//Program to create an embedded struct.

package main

import "fmt"

type Aurthor struct {
	AurthorName string
	AurthorAge  int
}

type Book struct {
	BookName string
	Aurthor  // Embedded struct
}

func main() {
	book1 := Book{
		BookName: "The Secret",
		Aurthor: Aurthor{
			AurthorName: "Rhonda Byrne",
			AurthorAge:  71,
		},
	}

	fmt.Println(book1)
	fmt.Println(book1.AurthorAge)
	fmt.Println(book1.AurthorName)
	fmt.Println(book1.Aurthor.AurthorAge)
}
