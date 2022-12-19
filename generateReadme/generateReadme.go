package main

//TODO: https://github.com/shurcooL/markdownfmt
import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

/*
generateReadme.go restores project structre and generates a readme file for all current solutions.
Any valid solution in the root folder will be moved to the designated folder

Definitions:
- _Solution file_ is a source code file that has a valid leetcode problem slug as a name plus an alowed extension.
- _Leetcode problem slug_ is a string that contains at least one hyphen and doesn't start with an undescore.
This definition for leetcode problem slug was made empirically.
- _Stranded solution -- file with a source code that resides in the root directory instead of designated one.

CW: This code was made with further expansion to other languages in mind. However
current project layout keeps it somewhat grounded, so adding another supproted
extension to the list of supportedExtensions will not be enough to truly support it.

Project structure:
root
    slug
        slug.go
    slug1
        slug1.go
    py
        slug.py
Additional files and folders will be ignored as long as their name is not a valid leetcode problem slug.

*/

var supportedExtensions = []string{
	".go",
	".py",
}

const PythonSolutionsFolder = "py"

var logger = log.New(os.Stderr, "", 0)

// Solution tuple that represents a leetcode problem solution
type Solution struct {
	slug string // leetcode problem slug
	path string // local path of solution
}

// isSolutionFile returns if the file is a valid solution file
func isSolutionFile(fileName string) bool {
	return contains(supportedExtensions, filepath.Ext(fileName)) &&
		isLeetCodeProblemSlug(getLeetcodeSlug(fileName))
}

// isLeetCodeProblemSlug returns if the string is a valid leetcode problem slug
func isLeetCodeProblemSlug(name string) bool {
	return !strings.HasPrefix(name, "_") && strings.ContainsRune(name, '-')
}

// getLeetcodeSlug returns a leetcode slug from a valid solution file
func getLeetcodeSlug(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// getDesignatedLocation returns folder where solution supposed to be stored and filepath to it
func getDesignatedLocation(fileName string) (string, string) {
	var folder string
	switch ext := filepath.Ext(fileName); ext {
	case ".py":
		folder = PythonSolutionsFolder
	case ".go":
		folder = getLeetcodeSlug(fileName)
	default:
		logger.Fatalln("Can't decide directory for", fileName)
	}
	return folder, filepath.Join(folder, fileName)
}

// getStrandedSolutions returns solutions in the root directory
func getStrandedSolutions(workDir string) (strandedSolutions []string) {
	files, err := os.ReadDir(workDir)
	if err != nil {
		logger.Fatalln(workDir, err)
	}
	var fileName string
	for _, file := range files {
		fileName = file.Name()
		if !file.IsDir() && isSolutionFile(fileName) {
			strandedSolutions = append(strandedSolutions, fileName)
		}
	}
	return
}

// organizeStrandedSolutions moves solutions from the root directory to designated folders
func organizeStrandedSolutions(workDir string) {
	strandedSolutions := getStrandedSolutions(workDir)
	var folderName, location string
	for _, fileName := range strandedSolutions {
		folderName, location = getDesignatedLocation(fileName)
		logger.Println("Moving", location)
		_ = os.Mkdir(folderName, fs.ModePerm)
		err := os.Rename(fileName, location)
		if err != nil {
			logger.Println(err)
			continue
		}
	}
}

// getSolutionFiles returns solutions for current project
func getSolutionFiles(workDir string) (solutions []Solution) {

	folders := map[string]string{
		".go": workDir,
		".py": filepath.Join(workDir, PythonSolutionsFolder),
	}
	for ext, folder := range folders {
		files, err := os.ReadDir(folder)
		if err != nil {
			logger.Fatalln(workDir, err)
		}
		for _, file := range files {
			fileName := file.Name()
			slug := getLeetcodeSlug(fileName)
			if isLeetCodeProblemSlug(slug) {
				_, location := getDesignatedLocation(slug + ext)
				solutions = append(solutions, Solution{slug, location})
			}
		}
	}
	return solutions
}

var checkOnlyFlag bool

func main() {
	var workDir string // working directory
	if workDir = flag.Arg(0); workDir == "" {
		workDir = "."
	}
	// Checking if there are stranded solutions present
	if checkOnlyFlag {
		os.Exit(len(getStrandedSolutions(workDir)))
	}
	organizeStrandedSolutions(workDir)
	solutions := getSolutionFiles(workDir)
	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i].slug < solutions[j].slug
	})
	fmt.Print(Header)
	for _, solution := range solutions {
		fmt.Printf(
			"|[%s](https://leetcode.com/problems/%s/)|[source code](%s)|\n",
			solution.slug,
			solution.slug,
			solution.path,
		)
	}
}

func init() {
	flag.BoolVar(&checkOnlyFlag, "check-only", false, "")
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

const Header = `
This file is autogenerated. Check [source code](generateReadme.go) or run [Makefile](Makefile)
if you wish to modify or regenerate its contents.

### Go solutions for leetcode

This project contains code from the Leetcode golang initiative. Trying to solve
leetcode problems simultaneously rewriting and porting a part of the python std
library without (sigh...) generics.
Largely hindered by the executive leetcode desicion to not update golang runtime
from 1.17 at the moment (feature request submitted).

[Check me out](https://leetcode.com/memosiki/) on leetcode!

| Problem | Solution |
| --- | --- |
`
