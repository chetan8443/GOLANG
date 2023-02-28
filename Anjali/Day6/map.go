package main

import "fmt"

func main() {
    // Declare a map of strings to integers
    m := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
    }

    // Create a slice to store the values of the map
    var values []int
    for _, value := range m {
        values = append(values, value)
    }

    // Print out the values in reverse order
    for i := len(values) - 1; i >= 0; i-- {
        fmt.Println(values[i])
    }
}
