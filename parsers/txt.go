package parsers

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// parseTxt takes a *.txt file and returns slices of each quote.
// It is expected that each new line in the txt file is a new quote.
func parseTxt(path string) []string {
	fmt.Println("parsing txt")

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open file " + path)
	}
	scanner := bufio.NewScanner(readFile)

	motds := make([]string, 0, 20)

	for scanner.Scan() {
		motds = append(motds, scanner.Text())
	}

	return motds
}
