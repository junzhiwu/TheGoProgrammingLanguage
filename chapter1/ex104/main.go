package main

//print the names of all files in which each duplicated line occurs.
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	sourceFiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, sourceFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, sourceFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s repeats %d times in %v\n", line, n, sourceFiles[line])
		}
	}

}

func countLines(f *os.File, counts map[string]int, sourceFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !existsInSourceFiles(f.Name(), sourceFiles[line]) {
			sourceFiles[line] = append(sourceFiles[line], f.Name())
		}
	}
}

func existsInSourceFiles(name string, fileNames []string) bool {
	for _, fileName := range fileNames {
		if name == fileName {
			return true
		}
	}
	return false
}
