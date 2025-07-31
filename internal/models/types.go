/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package models

import (
	"path/filepath"

	"github.com/Turtle-In-Space/theia/pkg/helpers"
)

// ----- Structs ----- //

// implemeted for future use
type Target struct {
	Name  string
	Hosts []Host
}

type Host struct {
	Hostname  string
	IPAddr    string
	Services  []Service
	DataDir   string
	ResultDir string
}

type Service struct {
	Name string
	Port int
}

func (t *Target) AddDirs(dataDir, resultDir string) {
	if len(t.Hosts) == 1 {
		host := &t.Hosts[0]
		host.DataDir = dataDir
		host.ResultDir = resultDir
	} else {
		for i := range t.Hosts {
			t.Hosts[i].addDirs(dataDir, resultDir)
		}
	}
}

// create a host and dirs for host
func (h *Host) addDirs(dataDir, resultDir string) {
	// create dirs for host
	h.DataDir = filepath.Join(dataDir, h.IPAddr)
	h.ResultDir = filepath.Join(resultDir, h.IPAddr)

	helpers.CreateDir(h.DataDir)
	helpers.CreateDir(h.ResultDir)
}
