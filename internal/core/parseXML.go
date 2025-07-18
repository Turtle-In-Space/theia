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
	Ports    xmlPorts     `xml:"ports"`
	Address  xmlAddress   `xml:"address"`
	Hostname xmlHostnames `xml:"hostnames"`
}

type xmlAddress struct {
	Addr string `xml:"addr,attr"`
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
	Protocol string     `xml:"protocol,attr"`
	PortID   int        `xml:"portid,attr"`
	Service  xmlService `xml:"service"`
}

type xmlService struct {
	Name string `xml:"name,attr"`
}

func GetTarget(xmlFilePath, targetName string) target {
	xmlFile := helpers.OpenFile(xmlFilePath)
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var results nmapRun
	xml.Unmarshal(byteValue, &results)

	return parseTarget(results, targetName)
}

// Stores all services in a slice
func parseTarget(results nmapRun, name string) target {
	return target{
		name:  name,
		hosts: parseHosts(results),
	}
}

// Stores all hosts in a slice
func parseHosts(results nmapRun) (hosts []host) {
	for _, newHost := range results.Hosts {

		//TODO: fix or explain this part
		hostnames := newHost.Hostname.Hostnames
		var name string
		if len(hostnames) == 0 {
			name = ""
		} else {
			name = hostnames[0].Name
		}

		hosts = append(hosts, host{
			hostname: name,
			ipAddr:   newHost.Address.Addr,
			services: parseServices(newHost),
		})
	}

	return
}

// Stores all services in a slice
func parseServices(newHost xmlHost) (services []service) {
	for _, port := range newHost.Ports.Ports {
		serviceName := port.Service.Name

		if serviceName == "" {
			serviceName = "unknown"
		}

		services = append(services, service{
			name:   serviceName,
			ipAddr: newHost.Address.Addr,
			port:   port.PortID,
		})
	}

	return
}
