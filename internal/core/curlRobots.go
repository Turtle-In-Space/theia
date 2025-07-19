/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"fmt"
	"os/exec"
)

// ----- Structs ----- //

type curlRobotsScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s curlRobotsScanner) Run(service service, host host) {
	resultFileName, dataFileName := fileNames(host, s.name, ".html", service.port)

	url := fmt.Sprintf("http://%s:%d/robots.txt", host.ipAddr, service.port)

	cmd := exec.Command("curl", url, "--fail", "--styled-output", "--output", resultFileName)
	execute(s, cmd, dataFileName)
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
	curlRobotScanner := curlRobotsScanner{
		serviceNames: []string{"http", "https"},
		name:         "curlRobotsScanner",
	}

	register("curlRobots", curlRobotScanner)
}
