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

	"github.com/Turtle-In-Space/theia/helpers"
	"github.com/Turtle-In-Space/theia/models"
	"github.com/Turtle-In-Space/theia/output"
)

// ----- Interfaces ----- //

type ServiceScanner interface {
	Run(port models.Port, host models.Host) (err error)
	ServiceNames() []string
	Name() string
}

// ----- Intensity ----- //

// set
// get
// access in core
/*
isValid()
scan.Intensity > Intensity
	false
*/

// ----- Variables ----- //

var serviceRegistry = make(map[string]ServiceScanner)

// ----- Public Functions ----- //

// get scanners that target a service
func ScannersByServiceName(service string) (scanners []ServiceScanner, ok bool) {
	for _, scanner := range serviceRegistry {
		if slices.Contains(scanner.ServiceNames(), service) {
			scanners = append(scanners, scanner)
		}
	}

	if len(scanners) == 0 {
		return nil, false
	}

	return scanners, true
}

// ----- Private Functions ----- //

// register the scanner to be used
func register(name string, scanner ServiceScanner) {
	serviceRegistry[name] = scanner
}

func execute(scanner ServiceScanner, cmd *exec.Cmd, resultFileName string) (exitCode int, err error) {
	_, err = exec.LookPath(cmd.Path)

	if err != nil {
		// cmd not found
		return -1, err
	}

	output.Info(output.Verbose, "Running %s", scanner.Name())
	output.Debug(cmd.String())

	if resultFileName != "" {
		resultFile := helpers.CreateFile(resultFileName)
		defer resultFile.Close()
		cmd.Stdout = resultFile
	}

	//TODO: move this to be cmd specific
	cmd.Env = append(cmd.Environ(), "NO_COLOR=1")

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode(), err
		}
		return 0, err
	}

	return 0, nil
}

// generate names for txt file and out file
func fileNames(scanName, dataExtension string, port models.Port) (resultFileName, dataFileName string) {
	result := fmt.Sprintf("%s.txt", scanName)
	data := fmt.Sprintf("%s%s", scanName, dataExtension)

	resultFileName = filepath.Join(port.Dir, result)
	dataFileName = filepath.Join(port.DataDir, data)

	return
}
