/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"os/exec"
)

type smbScanner struct {
	serviceNames []string
	name         string
}

// run the scan on a ipAddr for a port
func (s smbScanner) Run(service service, host host) {
	resultFileName, dataFileName := fileNames(s.name, host.ipAddr, service.port, ".txt", "")

	cmd := exec.Command("enum4linux-ng", "-A", host.ipAddr, "-oJ", dataFileName)
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

func init() {
	smbScanner := smbScanner{
		serviceNames: []string{"microsoft-ds", "netbios-ssn"},
		name:         "SMBScanner",
	}

	register("smb", smbScanner)
}
