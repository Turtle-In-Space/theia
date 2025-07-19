/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// ----- Structs ----- //

type webDirScanner struct {
	serviceNames []string
	name         string
}

// ----- Variables ----- //

var seclistPath string = filepath.Join("usr", "share", "seclists")

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s webDirScanner) Run(service service, host host) {
	resultFileName, dataFileName := fileNames(host, s.name, ".json", service.port)

	url := fmt.Sprintf("http://%s:%d", host.ipAddr, service.port)
	wordlist := filepath.Join(seclistPath, "Discovery", "Web-Content", "big.txt")

	cmd := exec.Command("ffuf", "-u", url, "-w", wordlist, "-ic", "-c", "-ach", "-o", dataFileName)
	execute(s, cmd, resultFileName)
}

// get all aliases for service names
func (s webDirScanner) ServiceNames() []string {
	return s.serviceNames
}

// get the name of this scanner
func (s webDirScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	webDirScanner := webDirScanner{
		serviceNames: []string{"http"},
		name:         "WebDirScanner",
	}

	register("webDir", webDirScanner)
}
