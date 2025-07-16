/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package helpers

import (
	"os"

	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// create a dir with the given name
func CreateDir(name string) {
	err := os.MkdirAll(name, 0766)

	if err != nil {
		out.Error("CreateDir: %s", err.Error())
	}
}

// open a file and handle errors
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		out.Error("OpenFile: %s", err.Error())
	}

	return file
}
