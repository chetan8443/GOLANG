package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    
    file, err := os.Open("file.json")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Read the JSON file into a byte slice
    slice, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    
    var data interface{}
    err = json.Unmarshal(slice, &data)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    // Print the data structure
    fmt.Println(data)
}
