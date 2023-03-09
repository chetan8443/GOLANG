package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

type Data struct {
    name string
    age  int
}

func main() {
    data, err := ioutil.ReadFile("./user.json")
    if err != nil {
        fmt.Print("Error ocuured:", err)
    }
    fmt.Println("the data is:", data)
    fmt.Print(string(data))
    var d1 Data
    err = json.Unmarshal(data, &d1)
    if err != nil {
        fmt.Print("Error ocuured:", err)
    }
    log.Printf("%s", d1.name)
    log.Printf("%v", d1.age)
}