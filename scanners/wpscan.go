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

type WPScan struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //
// --no-banner
// --url --output --format

func (s WPScan) Run(port models.Port, host models.Host) (err error) {
	resultFileName, _ := fileNames(s.name, "", port)

	url := createURL(host, port)
	cmd := exec.Command("wpscan", "--url", url)

	_, err = execute(s, cmd, resultFileName)
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s WPScan) ServiceNames() []string {
	return s.serviceNames
}

func (s WPScan) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := WPScan{
		serviceNames: []string{"http", "https"},
		name:         "WPScanner",
	}

	register("wpScanner", scanner)
}
