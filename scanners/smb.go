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

type smbScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s smbScanner) Run(port models.Port, host models.Host) (err error) {
	resultFileName, dataFileName := fileNames(s.name, "", port)

	cmd := exec.Command("enum4linux-ng", "-A", host.Address(), "-oJ", dataFileName)
	cmd.Env = append(cmd.Environ(), "NO_COLOR=1")

	_, err = execute(s, cmd, resultFileName)
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s smbScanner) ServiceNames() []string {
	return s.serviceNames
}

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
