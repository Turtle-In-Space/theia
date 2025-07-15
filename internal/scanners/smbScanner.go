/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

type SMBScanner struct {
	aliases []string
	name    string
}

// run the scan on a ipAddr for a port
func (s SMBScanner) Run(ipAddr string, port int) {
	out.Info("SMBScanner - %s:%d", ipAddr, port)
}

// get all aliases for service names
func (s SMBScanner) Aliases() (aliases []string) {
	return s.aliases
}

// get the name of this scanner
func (s SMBScanner) Name() (name string) {
	return s.name
}

func init() {
	smbScanner := SMBScanner{
		aliases: []string{"microsoft-ds", "netbios-ssn"},
		name:    "SMBScanner",
	}

	Register("smb", smbScanner)
}
