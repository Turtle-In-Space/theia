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

type SSHScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s SSHScanner) Run(port models.Port, host models.Host) (err error) {
	resultFileName, dataFileName := fileNames(s.name, ".xml", port)

	cmd := exec.Command(
		"nmap", "-Pn", "-T4", "-sV", "-p", port.ID,
		"-oN", resultFileName, "-oX", dataFileName,
		host.IPAddr)

	_, err = execute(s, cmd, "")
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

// get all aliases for service names
func (s SSHScanner) ServiceNames() []string {
	return s.serviceNames
}

// get the name of this scanner
func (s SSHScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := SSHScanner{
		serviceNames: []string{"ssh"},
		name:         "SSHScanner",
	}

	register("sshScanner", scanner)
}
