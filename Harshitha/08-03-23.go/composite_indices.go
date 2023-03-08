package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var Range int
    var size int
    fmt.Println("Enter the range :")
    fmt.Scanln(&Range)
    fmt.Println("Enter the number of random numbers you need:")
    fmt.Scanln(&size)
    var arr = make([]int, 0)
    for i := 0; i < size; i++ {
        arr = append(arr, rand.Intn(Range))
    }
	fmt.Println(arr)
	result:=make([]int,0)
	flag:=true
    for _, num := range arr {
        if num > 2 {
            for i := 2; i < num; i++ {
                if num%i == 0 {
                    flag=false
                } else {
                    continue
                }
            }
			if flag{
				result=duplicates(result,arr,num)
			}
        }
    }
	fmt.Println(result)
}
func duplicates(result []int,arr []int,num int)[]int{
	i:=0
	flag:=false
	for i=0;i<len(result);i++{
		if num!=result[i]{
			flag=true
		}

	}
	if flag{
		result=append(result,num)
	}
	return result
}