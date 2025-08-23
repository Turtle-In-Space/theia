/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"fmt"
	"os/exec"

	"github.com/Turtle-In-Space/theia/models"
)

// ----- Structs ----- //

type FTPScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s FTPScanner) Run(port models.Port, host models.Host) (err error) {
	resultFileName, dataFileName := fileNames(s.name, ".xml", port)

	cmd := exec.Command(
		"nmap", "-Pn", "-T4", "-sV", "-sC", "-p", port.ID,
		"-oN", resultFileName, "-oX", dataFileName,
		host.IPAddr)

	_, err = execute(s, cmd, "")
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s FTPScanner) ServiceNames() []string {
	return s.serviceNames
}

func (s FTPScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := FTPScanner{
		serviceNames: []string{"ftp"},
		name:         "FTPScanner",
	}

	register("ftpScanner", scanner)
}
