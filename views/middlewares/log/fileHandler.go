package log

import (
	"os"
)

var (
	newFile *os.File
	err     error
	path    string
)

// check file with given path is exists or not
func fileExists() bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return false
}

// create a new folder in given path
func createFile() bool {
	return false
}

func OpenFile(givenPath string) {
	path = givenPath

	if fileExists() {

	} else {

	}
}
