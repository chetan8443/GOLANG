package main
import "fmt"

func main(){
	arr:=[]int{5,8,5,7,8,9,20,7}
	var size int=len(arr)
	
	countFreq(arr,size)
}
func countFreq([] arr,size int){
	for i:=0;i<size;i++{
		flag:=0
		count:=0

		for j:=i+1;j<size;j++{
			if(arr[i]==arr[j]){
				flag=1
				break
			}
		}
		// if(flag==1)
		// continue

		for j:=0;j<=i;j++{
			if (arr[i] == arr[j])
                    count ++
		}

		fmt.Println(arr[i],":",count)
	}
}