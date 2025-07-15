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
	hosts []xmlHost `xml:"host"`
}

type xmlHost struct {
	ports xmlPorts `xml:"ports"`
}

type xmlPorts struct {
	ports []xmlPort `xml:"port"`
}

type xmlPort struct {
	protocol string     `xml:"protocol,attr"`
	portID   int        `xml:"portid,attr"`
	service  xmlService `xml:"service"`
}

type xmlService struct {
	name string `xml:"name,attr"`
}

// Parse all servies in `xmlFilePath`
// Returns map of [port]service
func GetServices(xmlFilePath string) []service {
	xmlFile := helpers.OpenFile(xmlFilePath)
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var results nmapRun
	xml.Unmarshal(byteValue, &results)

	return parseServices(results)
}

// Stores all services in a slice
func parseServices(results nmapRun) (services []service) {
	for _, host := range results.hosts {
		for _, port := range host.ports.ports {
			serviceName := port.service.name

			if serviceName == "" {
				serviceName = "unknown"
			}

			services = append(services, service{
				name: serviceName,
				port: port.portID,
			})
		}
	}

	return
}
