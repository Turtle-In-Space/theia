package core

import (
	"fmt"
	"path/filepath"

	"github.com/Turtle-In-Space/theia/helpers"
	"github.com/Turtle-In-Space/theia/models"
	"github.com/Turtle-In-Space/theia/output"
)

func CreateEnvFile(host models.Host) {
	var folderName string
	if host.Name != "" {
		folderName = host.Name
	} else {
		folderName = host.Address()
	}

	fileName := filepath.Join(host.Dir, fmt.Sprintf("../%s.env", folderName)) // bad fix?
	file := helpers.CreateFile(fileName)
	defer file.Close()

	helpers.WriteLine(file, "# Created by theia")
	helpers.WriteLine(file, fmt.Sprintf("ip=%s", host.IpAddr))

	if host.Hostname != "" {
		helpers.WriteLine(file, fmt.Sprintf("host=%s", host.Hostname))
	}

	helpers.WriteLine(file, fmt.Sprintf("url=http://%s", host.Address()))
	helpers.WriteLine(file, fmt.Sprintf("urls=https://%s", host.Address()))

	output.Debug("created env file for: %s", host.Address())
}
