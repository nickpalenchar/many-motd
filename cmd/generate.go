package cmd

import (
	"fmt"
	"log"
	"many-motd/config"
	"many-motd/parsers"
	"os"
)

func Generate() string {
	src := config.GetConfig().Src
	files := getSrcPaths(src)
	fmt.Printf("using path is %#v", files)

	motds := make([]string, 0, 30)

	for _, file := range files {
		motds = append(motds, parsers.Parse(file)...)
	}
	log.Printf("quotes received: %#v", motds)

	return "A very wise saynig"
}

func getSrcPaths(path string) []string {
	dir, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open directory " + path)
	}
	files, err := dir.Readdir(0)
	if err != nil {
		log.Fatal("Cannot open directory" + path)
	}

	filenames := make([]string, 0, 10)
	sep := string(os.PathSeparator)
	for _, val := range files {
		filenames = append(filenames, path+sep+val.Name())
	}
	return filenames[:]
}
