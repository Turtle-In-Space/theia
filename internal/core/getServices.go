/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

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

func GetServices(xmlFilePath string) []string {

	xmlFile, err := os.Open(xmlFilePath)

	if err != nil {
		log.Fatal(err)
	}

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var results NmapRun
	xml.Unmarshal(byteValue, &results)

	for _, host := range results.Hosts {
		for _, port := range host.Ports.Ports {

			var serviceName string
			if port.Service.Name == "" {
				serviceName = "unknown"
			} else {
				serviceName = port.Service.Name
			}
			fmt.Printf("Port: %d/%s - Service: %s\n", port.PortID, port.Protocol, serviceName)
		}
	}

	return nil
}
