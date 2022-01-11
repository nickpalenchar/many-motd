/*
Package parsers take various sources (usually filetypes)
and extracts slices of strings for use in motd.
*/
package parsers

import (
	"log"
	"many-motd/data"
	"path/filepath"
)

func Parse(path string) []data.Message {

	base := filepath.Base(path)

	if t, _ := filepath.Match(`*.txt`, base); t {
		return parseTxt(path)
	}
	if t, _ := filepath.Match(`github.json`, base); t {
		log.Printf("Found github file")
		return parseRemote(path)
	}
	motds := parseTxt(path)

	return motds[:]

}
