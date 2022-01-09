package formatters

import (
	"fmt"
	"strings"
)

func wrap(m string, w int) string {
	s := strings.Split(m, " ")
	lines := make([]string, 0)
	nextLine := make([]string, 0)
	count := 0
	for _, line := range s {
		// if the word itself is larger than the line, it must be split
		if count == 0 && len(line) > w {
			fmt.Printf("too big!")
			lines = append(lines, line[:w], line[w:])
			count = len(line[w:])
			continue
		}
		fmt.Printf("line before check %#v\n", lines)
		if count+len(line) > w-1 {
			fmt.Printf("line complete: %#v\n", strings.Join(nextLine, " "))
			fmt.Printf("line complete: %#v\n", append(lines, strings.Join(nextLine, " ")))
			lines = append(lines, strings.Join(nextLine, " "))
			nextLine = []string{line}
			count = len(line)
			fmt.Printf("line is now: %#v\n", lines)
			continue
		}
		nextLine = append(nextLine, line)
		count += len(line) + 1
		// lines = append(lines, strings.Join(nextLine, " "))
	}
	lines = append(lines, strings.Join(nextLine, " "))

	return strings.Join(lines, "\n")
}
