package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// readFile()
	// createFile()
	// readDirectory()
	recursiveReadDirectory()
}

func readFile() {
	file, err := os.Open("test.txt")
	if(err != nil) {
		return
	}
	defer file.Close()

	// get file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if(err != nil) {
		return
	}

	str := string(bs)
	fmt.Println(str)

	// alternative, faster way to read files
	bsutil, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	fmt.Println(string(bsutil))
}

func createFile() {
	file, err := os.Create("test.txt")
	if(err != nil) {
		return
	}
	defer file.Close()

	file.WriteString("test")
}

func readDirectory() {
	dir, err := os.Open(".")
	if(err != nil) {
		return
	}
	defer dir.Close()

	// Returns all file entries with parameter value of -1.
	fileInfos, err := dir.Readdir(-1)
	if(err != nil) {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}


}

// Reads all contents of directory and contents of subdirectories within.
func recursiveReadDirectory() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

