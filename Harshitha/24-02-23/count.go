package main

import ("fmt"
"net/http"
)

func main() {
	// var arr = []int32{-1, -2, 0, 1, 2}
	// plusMinus(arr)5
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8080",nil))

}
func handler(w http.ResponseWriter,r *http.Request){
	if r.Method=="POST"{
	plusMinus(w,r)
		}else {
			http.Error(w,"invalid method",http.StatusMethodNotAllowed)
			return
		}
}
func plusMinus(w http.ResponseWriter,r *http.Request){
	arr:=[]int{}
	a,err:=ioutil.ReadAll(r.Body)
err:=json.Unmarshal(a,&arr)
if err!=nil{
	http.Error(w,"invalid syntax",http.StatusBadRequest)
	jsonData,_:=json.Marshal(map[string]interface{}{"Message":"Invalid Syntax","Status Code":"400"})
w.Write(jsonData)
	return
}
	var cpos int = 0
	var cneg int = 0
	var czero int = 0
	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] > 0 {
			cpos = cpos + 1
		} else if arr[i] == 0 {
			czero = czero + 1
		} else {
			cneg = cneg + 1
		}
	}
	fmt.Printf("%.8f",float64(cpos) / float64(len(arr)))
	fmt.Println(cneg / len(arr))
	fmt.Println(czero / len(arr))
	jsonData,_:=json.Marshal(map[string]interface{}{"Message":"Counted Successfully!!","Status Code":"200 ok"})
}


// func plusMinus(arr []int32) {
	 // Write your code here
// 	var cpos int = 0
// 	var cneg int = 0
// 	var czero int = 0
// 	var i int
// 	for i = 0; i < len(arr); i++ {
// 		if arr[i] > 0 {
// 			cpos = cpos + 1
// 		} else if arr[i] == 0 {
// 			czero = czero + 1
// 		} else {
// 			cneg = cneg + 1
// 		}
// 	}
// 	fmt.Printf("%.8f",float64(cpos) / float64(len(arr)))
// 	fmt.Println(cneg / len(arr))
// 	fmt.Println(czero / len(arr))
// }

