/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package helpers

import (
	"log"
	"os"
)

func CreateDir(name string) {
	err := os.MkdirAll(name, 0766)

	if err != nil {
		log.Fatal(err)
	}
}

// open a file and handle errors
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
