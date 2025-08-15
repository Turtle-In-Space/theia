package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Turtle-In-Space/theia/pkg/helpers"
	"github.com/Turtle-In-Space/theia/pkg/output"
)

func CreateEnvFile(ipAddr, path string) {
	fileName := filepath.Join(path, fmt.Sprintf("%s.env", ipAddr))
	file := helpers.CreateFile(fileName)
	defer file.Close()

	writeLine(file, "# Created by theia")
	writeLine(file, fmt.Sprintf("ip=%s", ipAddr))

	output.Info(output.Verbose, "created env file for: %s", ipAddr)
}

func writeLine(file *os.File, msg string) {
	_, err := file.WriteString(msg + "\n")

	if err != nil {
		output.Warn(output.Normal, "failed writing to file: %s with error: %s", file.Name(), err.Error())
	}
}
