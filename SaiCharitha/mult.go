package main

import "fmt"

func main(){
    var n int
    fmt.Print("Enter any Integer Number : ")
    fmt.Scan(&n)
    i:=1
    /*     For loop as a Go's While     */
    for {
        if(i>10){ 
            break;
        }
        fmt.Println(n," X ",i," = ",n*i)
        i++
    }
}