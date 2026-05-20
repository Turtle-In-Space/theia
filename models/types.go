/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package models

import (
	"fmt"
	"path/filepath"

	"github.com/Turtle-In-Space/theia/helpers"
)

// ----- Structs ----- //

type Host struct {
	Name     string
	Hostname string
	IpAddr   string
	Ports    []Port
	Dir      string
}

type Port struct {
	ID       string
	Protocol string
	State    string //TODO: use this?
	Service  Service
	Dir      string
	DataDir  string
}

type Service struct {
	Name string
}

func (h *Host) AddDirs(dir string) {
	h.Dir = dir

	for i := range h.Ports {
		port := &h.Ports[i]

		port.Dir = filepath.Join(h.Dir, port.Name())
		port.DataDir = filepath.Join(port.Dir, "data")

		helpers.CreateDir(port.Dir)
		helpers.CreateDir(port.DataDir)
	}
}

func (h *Host) Address() (addr string) {
	if h.Hostname != "" {
		addr = h.Hostname
	} else {
		addr = h.IpAddr
	}

	return
}

func (p *Port) Name() string {
	return fmt.Sprintf("%s/%s", p.Protocol, p.ID)
}
