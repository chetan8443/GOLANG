package first

import "fmt"

func GetCountOfPair(arr [] int,sum int){
	count:=0
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			if(arr[i]+arr[j]==sum){
				count++
			}
		}
	}
	fmt.Println(count)
}