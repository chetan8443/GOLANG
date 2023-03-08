
package main
//Progam to create a text file and write content in it
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("execution started")
	var names =[]string{"vipul ","Yogesh ","Vivek ","Vishal "}
   file,err:=os.OpenFile("details.txt",os.O_WRONLY|os.O_CREATE,0644)// creating file and giving permissions to it
   defer file.Close()
   if err!=nil {
	log.Fatal(err)
   }
   bufferWriter:=bufio.NewWriter(file)    //using bufferwriter to write data in that file as string
	for _, v := range names {
		
		wr,err:=bufferWriter.WriteString(v)
		// 
		_=wr
		if err != nil {
			log.Fatal(err)
		}
		bufferWriter.Flush()    

	}
	
}
