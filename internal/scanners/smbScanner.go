/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"os/exec"

	out "github.com/Turtle-In-Space/theia/pkg/output"
)

type smbScanner struct {
	aliases []string
	name    string
}

// run the scan on a ipAddr for a port
func (s smbScanner) Run(ipAddr string, port int) {
	txtFile, outFile := fileNames(s.name, ipAddr, port, "")

	cmd := exec.Command("enum4linux-ng", "-A", ipAddr, "-oJ", outFile, ">", txtFile)
	err := cmd.Run()

	if err != nil {
		out.Error(err.Error())
	}
}

// get all aliases for service names
func (s smbScanner) Aliases() (aliases []string) {
	return s.aliases
}

// get the name of this scanner
func (s smbScanner) Name() (name string) {
	return s.name
}

func init() {
	smbScanner := smbScanner{
		aliases: []string{"microsoft-ds", "netbios-ssn"},
		name:    "SMBScanner",
	}

	Register("smb", smbScanner)
}
