package main

import (
    "fmt"
    "math"
)

func main() {
    a := math.Pow(2, 3)
    fmt.Println(a)
    var arm int
    fmt.Println("Enter number")
    fmt.Scanln(&arm)
    n := arm
    number := arm
    count := 0
    for arm > 0 {
        arm = arm / 10
        count = count + 1
    }
    fmt.Println(count)
    rem := 0
    sum := 0
    var c int = count
    for n > 0 {
        rem = n % 10
        sum = sum + int(math.Pow(float64(rem), float64(c)))
        n = n / 10
    }
    if sum == number {
        fmt.Println("The given number is an armstron number")
    } else {
        fmt.Println("Not an armstrong number")
    }

}