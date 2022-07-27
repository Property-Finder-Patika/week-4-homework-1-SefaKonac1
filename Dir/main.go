package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {

	path := "/"
	err := filepath.WalkDir(path, visitedFiles)
	fmt.Printf("err value %v\n", err)

}

/*this function lists files name in a specified directory*/
func visitedFiles(path string, di fs.DirEntry, err error) error {

	fmt.Println(path)
	return nil

}
