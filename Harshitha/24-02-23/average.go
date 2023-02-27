/*Write a program that declares an array of floats and calculates the average of the elements.*/
/*package main

import "fmt"

func main(){
	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scanln(&size)
	var array [5]float64
	fmt.Println("Enter the elements into the array")
	for i:=0;i<size;i++{
		fmt.Scanln(&array[i])
	}
	fmt.Println(array)
	var a int
	var result int
	a=suma(array [5]float64)
	result=a/len(array)
	fmt.Println(result)
}

func suma(array [5]float64) int{
	var total int=0
	var i int
	for i=0;i<size;i++{
		total=total+array[i]
	}
	return total
}*/
package main

import "fmt"

func main() {
    var n, i int
    fmt.Println("enter the size of the array")
    fmt.Scanln(&n)
    var arr [7]float64
    fmt.Print("Enter the items in array :=")
    for i = 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    var sum float64 = 0
    for i := 0; i <= n; i++ {
        sum += (arr[i])
    }
    avg := (float64(sum)) / (float64(n))
    fmt.Println("sum =", sum)
    fmt.Println("avg =", avg)
}