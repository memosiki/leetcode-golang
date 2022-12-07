package main

import (
    "flag"
    "fmt"
    "os"
)


// Solution files are files with go code, that have the name of 
// .
// This defini
func isSolutionFile(file os.File) bool {

}

func main() {
    var dirName string // working directory

    if dirName = flag.Arg(0); dirName == ""{
        dirName = "."
    }

    files, err := os.ReadDir(dirName)
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        info, _ := file.Info()
        fmt.Println(file.Name(), info.IsDir())
    }

}

func init(){
    flag.Parse()
}