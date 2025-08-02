/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"fmt"
	"os/exec"

	"github.com/Turtle-In-Space/theia/internal/models"
)

// ----- Structs ----- //

type curlRobotsScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s curlRobotsScanner) Run(service models.Service, host models.Host) {
	resultFileName, _ := fileNames(host, s.name, ".html", service.Port)

	url := fmt.Sprintf("http://%s:%d/robots.txt", host.IPAddr, service.Port)

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
