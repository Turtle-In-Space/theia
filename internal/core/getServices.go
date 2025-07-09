/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"encoding/xml"
	"io"

	helpers "github.com/Turtle-In-Space/theia/pkg"
)

// ----- Define xml sections ----- //
type NmapRun struct {
	Hosts []Host `xml:"host"`
}

type Host struct {
	Ports Ports `xml:"ports"`
}

type Ports struct {
	Ports []Port `xml:"port"`
}

type Port struct {
	Protocol string  `xml:"protocol,attr"`
	PortID   int     `xml:"portid,attr"`
	Service  Service `xml:"service"`
}

type Service struct {
	Name string `xml:"name,attr"`
}

// Parse all servies in `xmlFilePath`
// Returns map of [port]service
func GetServices(xmlFilePath string) map[int]string {
	xmlFile := helpers.OpenFile(xmlFilePath)
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var results NmapRun
	xml.Unmarshal(byteValue, &results)

	return parseServices(results)
}

// Stores all services in a map
// [PORT-ID] : [SERVICE-NAME]
func parseServices(results NmapRun) (services map[int]string) {
	services = make(map[int]string)

	for _, host := range results.Hosts {
		for _, port := range host.Ports.Ports {
			serviceName := port.Service.Name

			if serviceName == "" {
				serviceName = "unknown"
			}

			services[port.PortID] = serviceName
		}
	}

	return
}
