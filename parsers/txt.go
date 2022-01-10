package parsers

import (
	"bufio"
	"log"
	"many-motd/data"
	"os"
	"regexp"
)

// parseTxt takes a *.txt file and returns slices of each quote.
// It is expected that each new line in the txt file is a new quote.
func parseTxt(path string) []data.Message {
	log.Println("parsing txt")

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open file " + path)
	}
	scanner := bufio.NewScanner(readFile)

	motds := make([]data.Message, 0, 20)

	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			motds = append(motds, messageFromTxt(t))
		}
	}

	return motds
}

// messageFromTxt constructs a Message with an
// author if it is detected in the string
func messageFromTxt(s string) data.Message {
	re := regexp.MustCompile(`(.*)\((.*)\)$`)
	matches := re.FindAllStringSubmatch(s, -1)
	if len(matches) == 1 {
		return *data.NewMessage(matches[0][1], matches[0][2])
	}
	return *data.NewMessage(s, "")
}
