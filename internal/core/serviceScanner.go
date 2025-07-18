/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"fmt"
	"path/filepath"
	"slices"
)

type ServiceScanner interface {
	Run(service service, host host)
	ServiceNames() []string
	Name() string
}

var serviceRegistry = make(map[string]ServiceScanner)

// register the scanner to be used
func register(name string, scanner ServiceScanner) {
	serviceRegistry[name] = scanner
}

// get the correct scanner from a service name
func ScannerByServiceName(service string) (ServiceScanner, bool) {
	for _, scanner := range serviceRegistry {
		if slices.Contains(scanner.ServiceNames(), service) {
			return scanner, true
		}
	}

	return nil, false
}

// generate names for txt file and out file
func fileNames(scanName string, ipAddr string, port int, resultExtension string, dataExtension string) (resultFileName string, dataFileName string) {
	result := fmt.Sprintf("%d_%s%s", port, scanName, resultExtension)
	data := fmt.Sprintf("%d_%s%s", port, scanName, dataExtension)

	resultFileName = filepath.Join("results", ipAddr, result)
	dataFileName = filepath.Join("data", ipAddr, data)

	return
}
