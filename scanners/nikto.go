/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"os/exec"

	"github.com/Turtle-In-Space/theia/models"
)

// ----- Structs ----- //

type NiktoScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s NiktoScanner) Run(port models.Port, host models.Host) {
	resultFileName, _ := fileNames(s.name, ".txt", port)

	//FIX: sometimes returns error 1?
	cmd := exec.Command("nikto", "-host", host.IPAddr, "-port", port.ID, "-o", resultFileName)
	execute(s, cmd, "")
}

// get all aliases for service names
func (s NiktoScanner) ServiceNames() []string {
	return s.serviceNames
}

// get the name of this scanner
func (s NiktoScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := NiktoScanner{
		serviceNames: []string{"http"},
		name:         "NiktoScanner",
	}

	register("niktoSccanner", scanner)
}
