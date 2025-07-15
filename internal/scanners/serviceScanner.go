/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import "slices"

type ServiceScanner interface {
	Run(target string, port int)
	Aliases() []string
	Name() string
}

var ServiceRegistry = make(map[string]ServiceScanner)

func Register(name string, scanner ServiceScanner) {
	ServiceRegistry[name] = scanner
}

// get the correct scanner from a service name
func ScannerByServiceName(service string) (ServiceScanner, bool) {
	for _, scanner := range ServiceRegistry {
		if slices.Contains(scanner.Aliases(), service) {
			return scanner, true
		}
	}

	return nil, false
}
