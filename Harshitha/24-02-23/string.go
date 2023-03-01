package main

import "fmt"

func main(){
	str :="My name is Harshitha#Royal I am the $District Collector@of Kadapa1"
	var upper,lower,number,special,spaces int =0,0,0,0,0
	var i int=0
	for i=0;i<len(str);i++{
		if (str[i] >= 'A' && str[i] <= 'Z'){
            upper=upper+1
		}else if (str[i] >= 'a' && str[i] <= 'z'){
            lower=lower+1
		}else if (str[i]>= '0' && str[i]<= '9'){
            number=number+1
		}else if (str[i]==' '){
            spaces=spaces+1
		}else{
			special=special+1
		}
	}
	fmt.Println(upper,lower,number,special,spaces)
}