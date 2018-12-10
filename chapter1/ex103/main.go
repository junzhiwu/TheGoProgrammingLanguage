package main

//measure the difference in running time between different version of printing os.Args

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start)
	fmt.Println("for loop version1 costs: ", elapsed)

	start = time.Now()
	s1, sep1 := "", ""
	for _, arg := range os.Args {
		s1 += sep1 + arg
		sep1 = " "
	}
	fmt.Println(s1)
	elapsed = time.Since(start)
	fmt.Println("for loop version2 costs: ", elapsed)

	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	elapsed = time.Since(start)
	fmt.Println("strings.Join() way costs: ", elapsed)

	start = time.Now()
	fmt.Println(os.Args)
	elapsed = time.Since(start)
	fmt.Println("fmt.Println() direct print costs: ", elapsed)
}
