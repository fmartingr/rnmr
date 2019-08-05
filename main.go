package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultAllowedExtensions = "jpg,jpeg,png,gif,avi,mp4,mov"

var extensions = flag.String("extensions", defaultAllowedExtensions, "Comma separated extensions to allow")

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isValidFile(extension string) bool {
	allowedExtensions := strings.Split(*extensions, ",")
	return stringInSlice(extension[1:], allowedExtensions)
}

func readDir(path string, f os.FileInfo, err error) error {
	extension := filepath.Ext(f.Name())
	if !f.IsDir() && extension != "" && isValidFile(extension) {
		filenameDate := f.ModTime().Format("2006-01-02_15-04-05")
		newPath := filepath.Dir(path) + "/" + filenameDate + extension
		if path != newPath {
			fmt.Printf("%s => %s\n", path, newPath)
			err := os.Rename(path, newPath)
			if err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func main() {
	flag.Parse()

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Walking by %s...\n", path)

	err2 := filepath.Walk(path, readDir)
	if err2 != nil {
		panic(err2)
	}
}
