package main
//echo arguments with index (including os.Args[0])
import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
