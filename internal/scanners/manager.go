/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package scanners

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/Turtle-In-Space/theia/internal/models"
	"github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// ----- Interfaces ----- //

type ServiceScanner interface {
	Run(port models.Port, host models.Host)
	ServiceNames() []string
	Name() string
}

// ----- Variables ----- //

var serviceRegistry = make(map[string]ServiceScanner)

// ----- Public Functions ----- //

// get scanners that target a service
func ScannerByServiceName(service string) (scanners []ServiceScanner, ok bool) {
	for _, scanner := range serviceRegistry {
		if slices.Contains(scanner.ServiceNames(), service) {
			scanners = append(scanners, scanner)
		}
	}

	if len(scanners) > 0 {
		return scanners, true
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

	if resultFileName != "" {
		resultFile := helpers.CreateFile(resultFileName)
		defer resultFile.Close()
		cmd.Stdout = resultFile
	}

	cmd.Env = append(cmd.Environ(), "NO_COLOR=1")
	err = cmd.Run()

	if err != nil {
		out.Warn("command: %s - error: %s", cmd.String(), err.Error())
	}
}

// generate names for txt file and out file
func fileNames(scanName, dataExtension string, port models.Port) (resultFileName, dataFileName string) {
	result := fmt.Sprintf("%s.txt", scanName)
	data := fmt.Sprintf("%s%s", scanName, dataExtension)

	resultFileName = filepath.Join(port.Dir, result)
	dataFileName = filepath.Join(port.DataDir, data)

	return
}
