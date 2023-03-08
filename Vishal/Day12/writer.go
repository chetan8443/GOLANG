package main
//Progra to create a txt file and write selected names in that file based on comparison
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("execution started")//two lists first and second for writing names
	var firstList =[]string{"vipul ","Yogesh ","Vivek ","Vishal ","dhiraj","viraj ","prathmesh"}
	var secondList =[]string{"vipul ","Yogesh ","Vivek ","Vishal ","Arbind","viraj ","prathmesh "}
   file,err:=os.OpenFile("details.txt",os.O_WRONLY|os.O_CREATE,0644) // creating txt file and giving some permissions
   defer file.Close()
   if err!=nil {
	log.Fatal(err)
   }
   bufferWriter:=bufio.NewWriter(file)//using bufferwriter for writing strin that file
	for i, v := range firstList {
		if v==secondList[i] {
			
		
		wr,err:=bufferWriter.WriteString(v)   //
		// 
		_=wr
		if err != nil {
			log.Fatal(err)
		}
		bufferWriter.Flush()
	}
	}
	
}
