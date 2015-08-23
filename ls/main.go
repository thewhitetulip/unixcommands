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

	args := os.Args[1:]

	if len(args) == 0 {
		currentDirectory, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		listDirectory(currentDirectory)
	} else {

		folderFiles := make([]string, 0, 10)
		commands := make([]string, 0, 10)

		for _, value := range args {
			if strings.HasPrefix(value, "-") {
				commands = append(commands, value)
				continue
			}
			if Exists(value) {
				folderFiles = append(folderFiles, value)
				continue
			} else {
				fmt.Println("ls: no such file or directory ", value)
			}
		}
		numberOfDir := len(folderFiles)
		for _, value := range folderFiles {
			fileInfo, _ := os.Stat(value)

			if fileInfo.IsDir() {
				if numberOfDir > 0 {
					fmt.Println(value, ":")
				}
				listDirectory(value)
			} else {
				fmt.Println(value)
			}

		}
	}

}

func listDirectory(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(files)
	}

	for _, file := range files {
		tmpName := file.Name()
		if !strings.HasPrefix(tmpName, ".") {
			fmt.Printf("%s ", tmpName)
		}
	}
	fmt.Println()
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
