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

type CewlScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s CewlScanner) Run(port models.Port, host models.Host) (err error) {
	resultFileName, _ := fileNames(s.name, "", port)
	resultEmailFileName, _ := fileNames(fmt.Sprintf("%s-email", s.name), "", port)

	url := fmt.Sprintf("http://%s:%s", host.IpAddr, port.ID)

	cmd := exec.Command("cewl", url, "--depth=3", "--with-numbers", "--email", "-w", resultFileName, "--email_file", resultEmailFileName)

	_, err = execute(s, cmd, resultFileName)
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s CewlScanner) ServiceNames() []string {
	return s.serviceNames
}

func (s CewlScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := CewlScanner{
		serviceNames: []string{"http"},
		name:         "CewlScanner",
	}

	register("cewlScanner", scanner)
}
