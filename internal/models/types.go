/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package models

import (
	"fmt"
	"path/filepath"

	"github.com/Turtle-In-Space/theia/pkg/helpers"
)

// ----- Structs ----- //

// implemeted for future use
type Target struct {
	Name  string
	Hosts []Host
	Dir   string
}

type Host struct {
	Hostname string
	IPAddr   string
	Ports    []Port
	Dir      string
}

type Port struct {
	ID       string
	Protocol string
	State    string
	Service  Service
	Dir      string
	DataDir  string
}

type Service struct {
	Name string
}

func (t *Target) AddDirs(scanDir string) {
	t.Dir = scanDir

	if len(t.Hosts) == 1 {
		host := &t.Hosts[0]
		host.Dir = t.Dir
		host.addDirs()
	} else {
		for i := range t.Hosts {
			host := &t.Hosts[i]
			host.Dir = host.IPAddr
			host.addDirs()
		}
	}
}

func (h *Host) addDirs() {
	for i := range h.Ports {
		port := &h.Ports[i]

		port.Dir = filepath.Join(h.Dir, port.Name())
		port.DataDir = filepath.Join(port.Dir, "data")

		helpers.CreateDir(port.Dir)
		helpers.CreateDir(port.DataDir)
	}
}

func (p *Port) Name() string {
	return fmt.Sprintf("%s/%s", p.Protocol, p.ID)
}
