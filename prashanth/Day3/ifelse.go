package main

import (
  "fmt"
)

func main() {
 
    fmt.println("Understanding conditional statements")
  
  // problem statement:
  // Grade a student according to score as follows:
  // 0-40 -> F
  // 41-80 -> B
  // 81-90 ->A
  // 91+  -> A+
  
  score :=75
  if score >= 0 && score <= 40{
    fmt.println("F")
  }else if score >= 41 && score <= 80{
    fmt.println("B")
  }else if score >= 81 && score <= 90{
    fmt.println("A")
  }else if score >=91 {
    fmt.println("A+")
  }
}
