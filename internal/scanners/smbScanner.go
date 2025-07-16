/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"os"
	"os/exec"

	out "github.com/Turtle-In-Space/theia/pkg/output"
)

type smbScanner struct {
	aliases []string
	name    string
}

// run the scan on a ipAddr for a port
func (s smbScanner) Run(ipAddr string, port int) {
	txtFileName, outFileName := fileNames(s.name, ipAddr, port, "")

	txtFile, err := os.Create(txtFileName)
	if err != nil {
		out.Error("%s: create file error: %s", s.name, err.Error())
	}
	defer txtFile.Close()

	cmd := exec.Command("enum4linux-ng", "-A", ipAddr, "-oJ", outFileName)
	cmd.Stdout = txtFile

	out.Info("Running %s against %s", s.name, ipAddr)
	err = cmd.Run()
	if err != nil {
		out.Error("%s: command error: %s", s.name, err.Error())
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
