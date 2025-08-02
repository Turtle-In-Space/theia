/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package core

import (
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sync"

	"github.com/Turtle-In-Space/theia/internal/models"
	"github.com/Turtle-In-Space/theia/internal/scanners"
	"github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// TODO rename, add ipAddr
type validScanner struct {
	scanner scanners.ServiceScanner
	port    models.Port
	host    models.Host
}

var (
	scanDir string
	dataDir string
)

// ----- Public Functions ----- //

// begin the target scan
func ScanTarget(ip, targetName string) {
	createTargetStructure(targetName)
	dataOutPath := scanTarget(ip)
	target := GetTarget(dataOutPath, targetName)
	target.AddDirs(scanDir)

	scannerQueue := queueScanners(target)
	runScanners(scannerQueue)
}

// ----- Private Functions ----- //

func createTargetStructure(name string) {
	helpers.CreateDir(name)
	os.Chdir(name)

	scanDir = "scans"
	dataDir = filepath.Join(scanDir, "data")

	helpers.CreateDir("loot")
	helpers.CreateDir("exploits")
	helpers.CreateDir(scanDir)
	helpers.CreateDir(dataDir)
}

func scanTarget(ip string) (dataOut string) {
	dataOut = filepath.Join(dataDir, "nmap.xml")
	txtOut := filepath.Join(scanDir, "_nmap.txt")

	cmd := exec.Command("nmap", ip, "-oX", dataOut, "-oN", txtOut)
	err := cmd.Run()

	if err != nil {
		out.Error("scanAllPorts: %s", err.Error())
	}

	return
}

func queueScanners(target models.Target) (servicesWithScan []validScanner) {
	var foundScanners []string

	// find scan for each serivce
	for _, host := range target.Hosts {
		// clear scanners per host
		foundScanners = nil

		for _, port := range host.Ports {
			scanners, ok := scanners.ScannerByServiceName(port.Service.Name)

			if ok {
				for _, scan := range scanners {
					out.Info("Found service %s on port %d - using scan %s", port.Service.Name, port.Name(), scan.Name())
					if !slices.Contains(foundScanners, scan.Name()) {
						servicesWithScan = append(servicesWithScan,
							validScanner{
								scanner: scan,
								port:    port,
								host:    host,
							})
						foundScanners = append(foundScanners, scan.Name())
					}
				}
			} else {
				out.Warn("Found service %s on port %d - found no scan", port.Service.Name, port.Name())
			}
		}
	}

	return
}

// run all queued scanners and wait for them to finnish
func runScanners(scannerQueue []validScanner) {
	var wg sync.WaitGroup

	for _, scanner := range scannerQueue {
		wg.Add(1)
		go func() {
			defer wg.Done()
			scanner.scanner.Run(scanner.port, scanner.host)
		}()
	}

	wg.Wait()
}
