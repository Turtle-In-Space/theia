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

type NiktoScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s NiktoScanner) Run(port models.Port, host models.Host) (err error) {
	resultFileName, _ := fileNames(s.name, ".txt", port)

	cmd := exec.Command("nikto", "-host", host.IPAddr, "-port", port.ID, "-o", resultFileName)

	_, err = execute(s, cmd, "")
	if err != nil {
		//FIX: always returns error 1?
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s NiktoScanner) ServiceNames() []string {
	return s.serviceNames
}

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
