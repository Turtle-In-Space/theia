/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/Turtle-In-Space/theia/models"
)

// ----- Structs ----- //

type subDirScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s subDirScanner) Run(port models.Port, host models.Host) (err error) {
	if host.Hostname == "" {
		return fmt.Errorf("%s: hostname not set, wont run sub dir scan", s.Name())
	}

	_, dataFileName := fileNames(s.name, ".html", port)

	url := createURL(host, port)
	wordlist := filepath.Join(seclistPath, "Discovery", "DNS", "subdomains-top1million-110000.txt")
	hostFuzz := fmt.Sprintf("HOST: FUZZ.%s", host.Hostname)

	cmd := exec.Command("ffuf", "-u", url, "-w", wordlist, "-ic", "-noninteractive", "-t=100",
		"-H", hostFuzz, "-ac", "-v", "-of", "html", "-o", dataFileName)

	_, err = execute(s, cmd, "")
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s subDirScanner) ServiceNames() []string {
	return s.serviceNames
}

func (s subDirScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := subDirScanner{
		serviceNames: []string{"http", "https"},
		name:         "SubDirScanner",
	}

	register("subDir", scanner)
}
