package main

import (
  "fmt"
  "sort"
)

func main() {
  mymap := map[string]int{
    "Elliot": 25,
    "Sophie": 24, 
    "Fraser": 20,
  }
  
  // make an array of type string to store our keys 
  keys := []string{} 

  // iterate over the map and append all keys to our
  // string array of keys
  for key := range mymap {
    keys = append(keys, key)
  }

  // use the sort method to sort our keys array
  sort.Strings(keys)

  for _, key := range keys {
    fmt.Println(key, mymap[key])
  }
}
