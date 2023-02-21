package main

import "fmt"

       const name string = "Election commision India "

	const Year int = 2023
func main() {

	

	
	const status bool = false
    cname:=""
	cage:=0.0
	fmt.Print("ENTER YOUR NAME: ")
	fmt.Scanln(&cname)
	fmt.Print("ENTER YOUR AGE: ")
	fmt.Scanln(&cage)
	fmt.Println()
	fmt.Println()
	fmt.Println(name)
	fmt.Println("Year",Year)
	checkEligiility(cage,cname)

}

func checkEligiility(cage float64,cname string)  {
	const limit float64 = 18.0
	if(cage < limit){
		fmt.Println(cname,"YOU ARE NOT ELIGIBLE FOR VOTING")
		fmt.Printf(" Your age is %.2f , AGE MUST BE 18 OR MORE ",cage)
	}else{
		fmt.Println(cname," CONGRATULATIONS , YOU CAN VOTE THIS YEAR")
	}
}
