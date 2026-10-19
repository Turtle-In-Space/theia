/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package helpers

import (
	"os"

	"github.com/Turtle-In-Space/theia/output"
)

// ----- Public Functions ----- //

//TODO: Handle errors better

// create a dir with the given name
func CreateDir(name string) {
	err := os.MkdirAll(name, 0766)

	if err != nil {
		output.Error("CreateDir: %s", err.Error())
	}
}

// open a file and handle errors
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		output.Error("OpenFile: %s", err.Error())
	}

	return file
}

// create a file and handle errors
func CreateFile(filePath string) *os.File {
	file, err := os.Create(filePath)

	if err != nil {
		output.Error("CreateFile: %s", err.Error())
	}

	return file
}

func WriteLine(file *os.File, msg string) {
	_, err := file.WriteString(msg + "\n")

	if err != nil {
		output.Warn(output.Normal, "failed writing to file: %s with error: %s", file.Name(), err.Error())
	}
}
