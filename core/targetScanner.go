/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/Turtle-In-Space/theia/helpers"
	"github.com/Turtle-In-Space/theia/models"
	"github.com/Turtle-In-Space/theia/output"
	"github.com/Turtle-In-Space/theia/scanners"
)

// TODO: rename,
type validScanner struct {
	scanner scanners.ServiceScanner
	port    models.Port
	host    models.Host
}

var (
	scanDir string = "scans"
	dataDir string = filepath.Join(scanDir, "data")
)

// ----- Public Functions ----- //

func ScanTarget(host models.Host) (scanCount int, scanErrCount int) {
	createFileStructure(host.Name)
	dataOutPath := scanTarget(host.Address())

	ParseNmapResult(dataOutPath, &host)
	host.AddDirs(scanDir)
	CreateEnvFile(host)
	printFoundPorts(host)

	scannerQueue := queueScanners(host)
	scanCount, scanErrCount = runScanners(scannerQueue)

	return
}

// ----- Private Functions ----- //

func createFileStructure(name string) {
	output.Debug("Creating file stucture...")
	helpers.CreateDir(name)
	os.Chdir(name)

	helpers.CreateDir("loot")
	helpers.CreateDir("exploits")
	helpers.CreateDir(scanDir)
	helpers.CreateDir(dataDir)
}

// use nmap to scan the target for open ports
func scanTarget(ip string) (dataOut string) {
	output.Info(output.Normal, "Scanning target...")

	dataOut = filepath.Join(dataDir, "nmap.xml")
	txtOut := filepath.Join(scanDir, "_nmap.txt")

	cmd := exec.Command("nmap", "-sV", "-sC", "-T4", "-Pn", "-n", ip,
		"-oX", dataOut, "-oN", txtOut)

	err := cmd.Run()
	if err != nil {
		output.Error("scanTarget: %s", err.Error())
	}

	return
}

func printFoundPorts(host models.Host) {
	for _, port := range host.Ports {
		output.Info(output.Normal, "Found service %s on host %s port %s", port.Service.Name, host.Address(), port.Name())
	}
}

// for each port match a scan to the service
func queueScanners(host models.Host) (servicesWithScan []validScanner) {
	for _, port := range host.Ports {
		scanners, ok := scanners.ScannersByServiceName(port.Service.Name)

		//TODO: work out a solution for smb having same service multiple ports. Also http may exist on multiple ports same host
		if ok {
			scannerNames := make([]string, len(scanners))

			for i, scan := range scanners {
				servicesWithScan = append(servicesWithScan,
					validScanner{
						scanner: scan,
						port:    port,
						host:    host,
					})
				scannerNames[i] = scan.Name()
			}
			output.Info(output.Verbose, "For service %s on %s:%s - using following scans: %s",
				port.Service.Name, host.Address(), port.Name(), scannerNames)

		} else {
			output.Warn(output.Verbose, "For service %s on %s:%s - found no scan", port.Service.Name, host.Address(), port.Name())
		}
	}

	return
}

func runScanners(scannerQueue []validScanner) (int, int) {
	var wg sync.WaitGroup
	var errCount atomic.Int32

	for _, scanner := range scannerQueue {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := scanner.scanner.Run(scanner.port, scanner.host)
			if err != nil {
				errCount.Add(1)

				output.Warn(output.Verbose, err.Error())
			}
		}()
	}

	wg.Wait()

	return len(scannerQueue), int(errCount.Load())
}
