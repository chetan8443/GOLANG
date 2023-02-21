package main
import "fmt"

const a bool=true           //declaring & intitializing string variable
const b int=47            //declaring & intitializing string variable
const c string="CHETAN" //declaring & intitializing string variable
const d float64=3.14    //declaring & intitializing string variable

func main(){

  
    //printing all the constants ,After initializing constatnt variable we can't reinitialize them
    fmt.Printf("Boolean : %v \n integer: %d \n String : %s\n Float64 : %f", a,b,c,d)

    fun1()  //calling fun1

}

func fun1(){
	    const a =false
    	// const a=true  We can't redeclare constant variable in same block
	    fmt.Printf("function boolean : %v", a)
}
