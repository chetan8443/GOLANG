package main

import "fmt"

func main() {
	reverse()
}

func reverse() {
	var currency = map[string]float64{}

	currency["USD"] = 82.7294
	currency["Euro"] = 87.5821
	currency["British Pound"] = 99.6354
	currency["Swiss Franc"] = 88.292

	currency["Rupee"] = 1.00
	currency["Japanese Yen"] = 0.6068
	fmt.Println(currency)
	printAllMaps(currency)
}

func printAllMaps(ab map[string]float64) {
	fmt.Println("Currency rates \n")
	arr:=[]float64{}
	
	for k, v := range ab {
		fmt.Printf("Currency is %s and value is %f\n",k,v)
		arr=append(arr, v)
	}
 fmt.Println()
	for i := len(arr)-1; i >=0; i-- {
		fmt.Printf(" %f,",arr[i])
	}
}
