package main
 
import (
    "fmt"
	"log"
	"bufio"
	"os"
)
 
func main() {
    
	fmt.Print("Введите целое число: ")

	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

    if err !=nil {
        log.Fatal(err)
    }
    fmt.Printf("Вы ввели число: %s\n", str)
}