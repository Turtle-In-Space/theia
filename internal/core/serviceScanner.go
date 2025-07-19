/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// ----- Interfaces ----- //

type ServiceScanner interface {
	Run(service service, host host)
	ServiceNames() []string
	Name() string
}

// ----- Variables ----- //

var serviceRegistry = make(map[string]ServiceScanner)

// ----- Public Functions ----- //

// get the correct scanner from a service name
func ScannerByServiceName(service string) (ServiceScanner, bool) {
	for _, scanner := range serviceRegistry {
		if slices.Contains(scanner.ServiceNames(), service) {
			return scanner, true
		}
	}

	return nil, false
}

// ----- Private Functions ----- //

// register the scanner to be used
func register(name string, scanner ServiceScanner) {
	serviceRegistry[name] = scanner
}

func execute(scanner ServiceScanner, cmd *exec.Cmd, resultFileName string) {
	_, err := exec.LookPath(cmd.Path)

	if errors.Is(err, exec.ErrNotFound) {
		out.Warn("executable %s not found in $PATH, not running %s", cmd.Path, scanner.Name())
		return
	}

	out.Info("Running %s", scanner.Name())
	resultFile := helpers.CreateFile(resultFileName)
	cmd.Stdout = resultFile
	// err = cmd.Run()
	//
	// if err != nil {
	// 	out.Error("%s: command error: %s", scanner.Name(), err.Error())
	// }
}

// generate names for txt file and out file
func fileNames(host host, scanName, dataExtension string, port int) (resultFileName, dataFileName string) {
	result := fmt.Sprintf("%d_%s.txt", port, scanName)
	data := fmt.Sprintf("%d_%s%s", port, scanName, dataExtension)

	resultFileName = filepath.Join(host.resultDir, result)
	dataFileName = filepath.Join(host.dataDir, data)

	return
}
