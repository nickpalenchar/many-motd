package parsers

import (
	"bufio"
	"log"
	"os"
)

// parseTxt takes a *.txt file and returns slices of each quote.
// It is expected that each new line in the txt file is a new quote.
func parseTxt(path string) []string {
	log.Println("parsing txt")

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open file " + path)
	}
	scanner := bufio.NewScanner(readFile)

	motds := make([]string, 0, 20)

	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			motds = append(motds, t)
		}
	}

	return motds
}
