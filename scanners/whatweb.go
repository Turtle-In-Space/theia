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

type WhatWebScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

func (s WhatWebScanner) Run(port models.Port, host models.Host) (err error) {
	resultFileName, _ := fileNames(s.name, "", port)

	url := createURL(host, port)
	cmd := exec.Command("whatweb", url, "--color=never") // add -v

	_, err = execute(s, cmd, resultFileName)
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name(), err)
	}

	return nil
}

func (s WhatWebScanner) ServiceNames() []string {
	return s.serviceNames
}

func (s WhatWebScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	scanner := WhatWebScanner{
		serviceNames: []string{"http", "https"},
		name:         "WhatWebScanner",
	}

	register("whatwebScanner", scanner)
}
