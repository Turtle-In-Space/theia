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
func GetTarget(xmlFilePath, targetName string) models.Target {
	output.Debug("Parsing nmap xml data...")

	xmlFile := helpers.OpenFile(xmlFilePath)
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var results nmapRun
	xml.Unmarshal(byteValue, &results)

	return parseTarget(results, targetName)
}

// ----- Private Functions ----- //

func parseTarget(results nmapRun, name string) models.Target {
	return models.Target{
		Name:  name,
		Hosts: parseHosts(results),
	}
}

func parseHosts(results nmapRun) (hosts []models.Host) {
	for _, newHost := range results.Hosts {
		var name, ipAddr string

		//TODO: sometimes multiple hostnames for some reason, use first as all are the same?
		hostnames := newHost.Hostname.Hostnames
		if len(hostnames) == 0 {
			name = ""
		} else {
			name = hostnames[0].Name
		}

		//TODO: return mac, v4 and v6?
		for _, addr := range newHost.Addresses {
			if addr.Type == "ipv4" {
				ipAddr = addr.Addr
				break
			}
		}

		hosts = append(hosts, models.Host{
			Hostname: name,
			IPAddr:   ipAddr,
			Ports:    parsePorts(newHost),
		})
	}

	return
}

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
