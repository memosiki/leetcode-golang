package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var supportedExtensions = []string{
	".go",
	".py",
}

var loggerInfo = log.New(os.Stdout, "INFO: ", 0)
var loggerErr = log.New(os.Stderr, "ERR: ", 0)

// isSolutionFile returns if the file is solution file or not
//
// Solution file is a source code file that has a valid leetcode problem slug as a name plus an alowed extension.
// Leetcode problem slug is a string that contains at least one hyphen and doesn't start with an undescore.
// This definition for leetcode problem slug was made empirically.
func isSolutionFile(fileName string) bool {
	return contains(supportedExtensions, filepath.Ext(fileName)) &&
		!strings.HasPrefix(fileName, "_") &&
		strings.ContainsRune(fileName, '-')
}

// getLeetcodeSlug returns a leetcode slug from a valid solution file
func getLeetcodeSlug(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func main() {
	var dirName string // working directory
	var err error
	if dirName = flag.Arg(0); dirName == "" {
		dirName = "."
	}

	files, err := os.ReadDir(dirName)
	if err != nil {
		loggerErr.Fatalln(dirName, err)
	}
	var fileName, folderName string
	for _, file := range files {
		fileName = file.Name()
		if !file.IsDir() && isSolutionFile(fileName) {
			folderName = getLeetcodeSlug(fileName)
			newLoc := filepath.Join(folderName, fileName)
			loggerInfo.Println("Moving solution to", newLoc)
			err = os.Mkdir(folderName, fs.ModePerm)
			if err != nil {
				loggerErr.Println(err)
				continue
			}
			err = os.Rename(fileName, newLoc)
			if err != nil {
				loggerErr.Println(err)
			}
		}
	}

}

func init() {
	flag.Parse()
}

func contains(container []string, elem string) bool {
	for _, val := range container {
		if val == elem {
			return true
		}
	}
	return false
}
