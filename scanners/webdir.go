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

type webDirScanner struct {
	serviceNames []string
	name         string
}

// ----- Variables ----- //

var seclistPath string = filepath.Join("/usr", "share", "SecLists")

// ----- Public Functions ----- //

func (s webDirScanner) Run(port models.Port, host models.Host) (err error) {
	_, dataFileName := fileNames(s.name, ".html", port)

	url := fmt.Sprintf("http://%s:%s/FUZZ", host.Address(), port.ID)
	wordlist := filepath.Join(seclistPath, "Discovery", "Web-Content", "big.txt")

	cmd := exec.Command("ffuf", "-u", url, "-w", wordlist, "-ic", "-noninteractive",
		"-ac", "-v", "-of", "html", "-o", dataFileName)

	_, err = execute(s, cmd, "")
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s webDirScanner) ServiceNames() []string {
	return s.serviceNames
}

func (s webDirScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := webDirScanner{
		serviceNames: []string{"http"},
		name:         "WebDirScanner",
	}

	register("webDir", scanner)
}
