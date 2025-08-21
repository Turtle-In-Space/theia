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

type curlRobotsScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s curlRobotsScanner) Run(port models.Port, host models.Host) {
	resultFileName, _ := fileNames(s.name, ".html", port)

	url := fmt.Sprintf("http://%s:%s/robots.txt", host.IPAddr, port.ID)

	cmd := exec.Command("curl", url, "--fail", "--styled-output", "--output", resultFileName)
	execute(s, cmd, "")
}

// get all aliases for service names
func (s curlRobotsScanner) ServiceNames() []string {
	return s.serviceNames
}

// get the name of this scanner
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
