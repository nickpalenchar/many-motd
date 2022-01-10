/*
Package parsers take various sources (usually filetypes)
and extracts slices of strings for use in motd.
*/
package parsers

import "many-motd/data"

func Parse(path string) []data.Message {

	motds := parseTxt(path)

	return motds[:]

}
