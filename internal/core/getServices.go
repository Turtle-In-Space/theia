/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"encoding/xml"
	"io"

	helpers "github.com/Turtle-In-Space/theia/pkg/helpers"
)

// ----- Define xml sections ----- //
type nmapRun struct {
	Hosts []xmlHost `xml:"host"`
}

type xmlHost struct {
	Ports xmlPorts `xml:"ports"`
}

type xmlPorts struct {
	Ports []xmlPort `xml:"port"`
}

type xmlPort struct {
	Protocol string     `xml:"protocol,attr"`
	PortID   int        `xml:"portid,attr"`
	Service  xmlService `xml:"service"`
}

type xmlService struct {
	Name string `xml:"name,attr"`
}

// Parse all servies in `xmlFilePath`
// Returns map of [port]service
func GetServices(xmlFilePath string) map[int]string {
	xmlFile := helpers.OpenFile(xmlFilePath)
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var results nmapRun
	xml.Unmarshal(byteValue, &results)

	return parseServices(results)
}

// Stores all services in a map
// [PORT-ID] : [SERVICE-NAME]
func parseServices(results nmapRun) (services map[int]string) {
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
