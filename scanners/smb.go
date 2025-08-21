/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"os/exec"

	"github.com/Turtle-In-Space/theia/models"
)

// ----- Structs ----- //

type smbScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s smbScanner) Run(port models.Port, host models.Host) {
	resultFileName, dataFileName := fileNames(s.name, "", port)

	cmd := exec.Command("enum4linux-ng", "-A", host.IPAddr, "-oJ", dataFileName)
	execute(s, cmd, resultFileName)
}

// get all aliases for service names
func (s smbScanner) ServiceNames() []string {
	return s.serviceNames
}

// get the name of this scanner
func (s smbScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := smbScanner{
		serviceNames: []string{"microsoft-ds", "netbios-ssn"},
		name:         "SMBScanner",
	}

	register("smb", scanner)
}
