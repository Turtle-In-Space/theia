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

type WhatWebScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s WhatWebScanner) Run(service models.Service, host models.Host) {
	resultFileName, _ := fileNames(host, s.name, "", service.Port)

	url := fmt.Sprintf("http://%s:%d", host.IPAddr, service.Port)

	cmd := exec.Command("whatweb", url, "-v")
	execute(s, cmd, resultFileName)
}

// get all aliases for service names
func (s WhatWebScanner) ServiceNames() []string {
	return s.serviceNames
}

// get the name of this scanner
func (s WhatWebScanner) Name() (name string) {
	return s.name
}

// ----- Private Functions ----- //

func init() {
	whatwebScanner := WhatWebScanner{
		serviceNames: []string{"http"},
		name:         "WhatWebScanner",
	}

	register("whatwebScanner", whatwebScanner)
}
