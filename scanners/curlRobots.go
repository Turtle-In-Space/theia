/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Turtle-In-Space/theia/models"
)

// ----- Structs ----- //

type curlRobotsScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s curlRobotsScanner) Run(port models.Port, host models.Host) (err error) {
	resultFileName, _ := fileNames(s.name, ".html", port)

	url := fmt.Sprintf("http://%s:%s/robots.txt", host.Address(), port.ID)

	cmd := exec.Command("curl", url, "--fail", "--styled-output", "--output", resultFileName)

	exitCode, err := execute(s, cmd, "")

	if err != nil {
		switch exitCode {
		case 22:
			// page not found - no error
			os.Remove(resultFileName)
		default:
			return fmt.Errorf("%s: %w", s.Name(), err)
		}
	}

	return nil
}

func (s curlRobotsScanner) ServiceNames() []string {
	return s.serviceNames
}

func (s curlRobotsScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := curlRobotsScanner{
		serviceNames: []string{"http", "https"},
		name:         "curlRobotsScanner",
	}

	register("curlRobots", scanner)
}
