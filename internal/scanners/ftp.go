/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"os/exec"

	"github.com/Turtle-In-Space/theia/internal/models"
)

// ----- Structs ----- //

type FTPScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s FTPScanner) Run(port models.Port, host models.Host) {
	resultFileName, dataFileName := fileNames(s.name, ".xml", port)

	cmd := exec.Command(
		"nmap", "-Pn", "-T4", "-sV", "-sC", "-p", port.ID,
		"-oN", resultFileName, "-oX", dataFileName,
		host.IPAddr)

	execute(s, cmd, "")
}

// get all aliases for service names
func (s FTPScanner) ServiceNames() []string {
	return s.serviceNames
}

// get the name of this scanner
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
