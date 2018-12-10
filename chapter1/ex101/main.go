package main

//print command-line arguments (including os.Args[0]
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// use for loop with index
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// use short variable declaration, range and blank identifier
	s1, sep1 := "", ""
	for _, arg := range os.Args {
		s1 += sep1 + arg
		sep1 = " "
	}
	fmt.Println(s1)
	// use Join function from strings package
	fmt.Println(strings.Join(os.Args, " "))
	// directly print the arguments
	fmt.Println(os.Args)
}
