package main
import (
    "fmt"
    "io/ioutil"
    "log"
	"os"
)
func main() {
	var filename string
	fmt.Scan(&filename)
	file,err := os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
    b, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(b))
}
