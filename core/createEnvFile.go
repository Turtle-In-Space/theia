package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Turtle-In-Space/theia/helpers"
	"github.com/Turtle-In-Space/theia/models"
	"github.com/Turtle-In-Space/theia/output"
)

func CreateEnvFile(host models.Host) {
	fileName := filepath.Join(host.Dir, fmt.Sprintf("%s.env", host.IpAddr))
	file := helpers.CreateFile(fileName)
	defer file.Close()

	writeLine(file, "# Created by theia")
	writeLine(file, fmt.Sprintf("ip=%s", host.IpAddr))

	if host.Hostname != "" {
		writeLine(file, fmt.Sprintf("host=%s", host.Hostname))
		writeLine(file, fmt.Sprintf("url=http://%s", host.Hostname))
		writeLine(file, fmt.Sprintf("urls=https://%s", host.Hostname))
	}

	output.Debug("created env file for: %s", host.IpAddr)
}

func writeLine(file *os.File, msg string) {
	_, err := file.WriteString(msg + "\n")

	if err != nil {
		output.Warn(output.Normal, "failed writing to file: %s with error: %s", file.Name(), err.Error())
	}
}
