/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/Turtle-In-Space/theia/helpers"
	"github.com/Turtle-In-Space/theia/models"
	"github.com/Turtle-In-Space/theia/output"
)

// ----- Structs ----- //
type nmapRun struct {
	Hosts []xmlHost `xml:"host"`
}

type xmlHost struct {
	Ports     xmlPorts     `xml:"ports"`
	Addresses []xmlAddress `xml:"address"`
	Hostname  xmlHostnames `xml:"hostnames"`
}

type xmlAddress struct {
	Addr string `xml:"addr,attr"`
	Type string `xml:"addrtype,attr"`
}

type xmlHostnames struct {
	Hostnames []xmlHostname
}

type xmlHostname struct {
	Name string `xml:"name,attr"`
}

type xmlPorts struct {
	Ports []xmlPort `xml:"port"`
}

type xmlPort struct {
	ID       int        `xml:"portid,attr"`
	Protocol string     `xml:"protocol,attr"`
	State    xmlState   `xml:"state"`
	Service  xmlService `xml:"service"`
}

type xmlState struct {
	State string `xml:"state,attr"`
}

type xmlService struct {
	Name string `xml:"name,attr"`
}

// ----- Public Functions ----- //

// parse the target from a nmap scan
func ParseNmapResult(xmlFilePath string, host *models.Host) {
	output.Debug("Parsing nmap xml data...")

	xmlFile := helpers.OpenFile(xmlFilePath)
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var results nmapRun
	xml.Unmarshal(byteValue, &results)
	host.Ports = parsePorts(results.Hosts[0])
}

// ----- Private Functions ----- //
func parsePorts(newHost xmlHost) (ports []models.Port) {
	for _, port := range newHost.Ports.Ports {

		ports = append(ports, models.Port{
			ID:       fmt.Sprintf("%d", port.ID),
			Protocol: port.Protocol,
			State:    port.State.State,
			Service:  parseService(port),
		})
	}

	return
}

func parseService(port xmlPort) models.Service {
	serviceName := port.Service.Name

	if serviceName == "" {
		serviceName = "unknown"
	}

	return models.Service{
		Name: serviceName,
	}
}
