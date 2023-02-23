/*

Problem
In a classic chase, Tom is running after Jerry as Jerry has eaten Tom's favourite food.

Jerry is running at a speed of 
�
X metres per second while Tom is chasing him at a speed of 
�
Y metres per second. Determine whether Tom will be able to catch Jerry.

Note that initially Jerry is not at the same position as Tom.

Input Format
The first line of input will contain a single integer 
�
T, denoting the number of test cases.
Each test case consists of two space-separated integers 
�
X and 
�
Y — the speeds of Jerry and Tom respectively.
Output Format
For each test case, output on a new line, YES, if Tom will be able to catch Jerry. Otherwise, output NO.

You can print each character in uppercase or lowercase. For example NO, no, No, and nO are all considered the same.

*/


package main


import "fmt"

func main() {
    var t int
    fmt.Scanln(&t)

    for i := 0; i < t; i++ {
        var x, y int
        fmt.Scanf("%d %d",&x,&y)
        
        if x <= y {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}