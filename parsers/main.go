// Package parsers take various sources (usually filetypes)
// and extracts slices of strings for use in motd.

package parsers

func Parse(path string) []string {

	motds := parseTxt(path)

	return motds[:]

}
