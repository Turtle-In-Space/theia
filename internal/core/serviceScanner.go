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
	Run(target string, port int)
	ServiceNames() []string
	Name() string
}

var serviceRegistry = make(map[string]ServiceScanner)

// register the scanner to be used
func Register(name string, scanner ServiceScanner) {
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
func fileNames(scanName string, ipAddr string, port int, resultExtension string, outExtension string) (resultFile string, outFile string) {
	txt := fmt.Sprintf("%d_%s%s", port, scanName, resultExtension)
	out := fmt.Sprintf("%d_%s%s", port, scanName, outExtension)

	resultFile = filepath.Join("results", ipAddr, txt)
	outFile = filepath.Join("xml", ipAddr, out)

	return
}
