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
		log.Panic(err)
	}
}
