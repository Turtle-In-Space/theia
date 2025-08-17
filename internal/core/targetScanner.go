/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/Turtle-In-Space/theia/internal/models"
	"github.com/Turtle-In-Space/theia/internal/scanners"
	"github.com/Turtle-In-Space/theia/pkg/helpers"
	"github.com/Turtle-In-Space/theia/pkg/output"
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

// begin the target scan
func ScanTarget(ip, targetName string) (scanCount int) {
	createFileStructure(targetName)
	dataOutPath := scanTarget(ip)

	target := GetTarget(dataOutPath, targetName)
	target.AddDirs(scanDir)
	addEnvFiles(target)
	printFoundPorts(target)

	scannerQueue := queueScanners(target)
	scanCount = runScanners(scannerQueue)

	return
}

// ----- Private Functions ----- //

// create the stucture
func createFileStructure(name string) {
	output.Debug("Creating file stucture...")
	helpers.CreateDir(name)
	os.Chdir(name)

	helpers.CreateDir("loot")
	helpers.CreateDir("exploits")
	helpers.CreateDir(scanDir)
	helpers.CreateDir(dataDir)
}

// use nmap to scan the target
func scanTarget(ip string) (dataOut string) {
	output.Info(output.Normal, "Scanning target...")

	dataOut = filepath.Join(dataDir, "nmap.xml")
	txtOut := filepath.Join(scanDir, "_nmap.txt")

	//TODO: remove comment
	cmd := exec.Command("nmap" /*,  "-sV" */, "-T4", "-Pn", ip,
		"-oX", dataOut, "-oN", txtOut)

	err := cmd.Run()
	if err != nil {
		output.Error("scanTarget: %s", err.Error())
	}

	return
}

// print all found ports for the target
func printFoundPorts(target models.Target) {
	for _, host := range target.Hosts {
		for _, port := range host.Ports {
			output.Info(output.Normal, "Found service %s on host %s port %s", port.Service.Name, host.IPAddr, port.Name())
		}
	}
}

// for each port match a scan to the service
func queueScanners(target models.Target) (servicesWithScan []validScanner) {
	for _, host := range target.Hosts {
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
					port.Service.Name, host.IPAddr, port.Name(), scannerNames)

			} else {
				output.Warn(output.Verbose, "For service %s on %s:%s - found no scan", port.Service.Name, host.IPAddr, port.Name())
			}
		}
	}

	return
}

// run all queued scanners and wait for them to finish, return count of scanners
func runScanners(scannerQueue []validScanner) (scanCount int) {
	var wg sync.WaitGroup

	for _, scanner := range scannerQueue {
		wg.Add(1)
		go func() {
			defer wg.Done()
			scanner.scanner.Run(scanner.port, scanner.host)
		}()
	}

	wg.Wait()

	return len(scannerQueue)
}

// create env files for each host in target
func addEnvFiles(target models.Target) {
	for _, host := range target.Hosts {
		CreateEnvFile(host.IPAddr, host.Dir)
	}
}
