// Pgm to check if 2 strings are anagrams(equal no. of characters, only position is different, for eg : silent,listen)
package main

import (
	"fmt"
	"sort"
	"strings"
)

//main function
func main() {
	var str1, str2 string
	fmt.Println("Enter 2 strings")
	fmt.Scan(&str1, &str2)
	anagram(str1, str2)
}

//Function to check if 2 strings are anagrams 
func anagram(st1 string, st2 string) {
	l1 := len(st1)
	l2 := len(st2)
	var isAnagram byte

	//if lengths are different , the strings cannot be anagrams
	if l1 != l2 {
		fmt.Println(st1, st2, "are not anagrams")
	}

	//converting each string into a slice and using the sorting method in ascending order
	st1_o := strings.Split(st1,"")
	sort.Strings(st1_o)
	strings.Join(st1_o,"")
	st2_o := strings.Split(st2,"")
	sort.Strings(st2_o)
	strings.Join(st2_o,"")

	//checking each letter in the sorted slice for equality
	for i := 0; i < l1; i++ {
		if st1_o[i] == st2_o[i]{
			isAnagram = 'Y'
		}else{
			isAnagram = 'N'
			break
		}	
	}
	if isAnagram == 'Y'{
		fmt.Println(st1,st2,"are anagrams")
	}else{
		fmt.Println(st1,st2,"are not anagrams")
	}
}
