//This is a go port of the standard ls command of the *nix OSes,
//currently it is basic, will become something great one day
//author: Suraj Patil date: 22 August 2015

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fileInfo, err := ioutil.ReadDir(currentDirectory)

	for _, file := range fileInfo {
		tmpName := file.Name()
		if !strings.HasPrefix(tmpName, ".") {
			fmt.Printf("%s ", tmpName)
		}
	}
	fmt.Println()
}
