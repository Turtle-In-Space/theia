/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"fmt"
	"os/exec"
)

// ----- Structs ----- //

type WhatWebScanner struct {
	serviceNames []string
	name         string
}

// ----- Public Functions ----- //

// run the scan on a ipAddr for a port
func (s WhatWebScanner) Run(service service, host host) {
	resultFileName, _ := fileNames(host, s.name, "", service.port)

	url := fmt.Sprintf("http://%s:%d", host.ipAddr, service.port)

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
