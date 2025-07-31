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

	models "github.com/Turtle-In-Space/theia/internal/models"
	scanners "github.com/Turtle-In-Space/theia/internal/scanners"
	"github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// TODO rename, add ipAddr
type validScanner struct {
	scanner scanners.ServiceScanner
	service models.Service
	host    models.Host
}

// ----- Variables ----- //

var (
	dataDir   string
	resultDir string
)

// ----- Public Functions ----- //

// begin the target scan
func ScanTarget(ip, targetName string) {
	createTargetStructure(targetName)
	dataOutPath := scanTarget(ip)
	target := GetTarget(dataOutPath, targetName)
	target.AddDirs(dataDir, resultDir)

	scannerQueue := queueScanners(target)
	runScanners(scannerQueue)
}

// ----- Private Functions ----- //

func createTargetStructure(name string) {
	helpers.CreateDir(name)
	os.Chdir(name)

	dataDir = filepath.Join("scans", "data/")
	resultDir = filepath.Join("scans", "results/")

	helpers.CreateDir("loot")
	helpers.CreateDir("exploits")
	helpers.CreateDir("scans")
	helpers.CreateDir(dataDir)
	helpers.CreateDir(resultDir)
}

func scanTarget(ip string) (dataOut string) {
	dataOut = filepath.Join(dataDir, "ports.xml")
	txtOut := filepath.Join(resultDir, "ports.txt")

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

		for _, service := range host.Services {
			scanners, ok := scanners.ScannerByServiceName(service.Name)

			if ok {
				for _, scan := range scanners {
					out.Info("Found service %s on port %d - using scan %s", service.Name, service.Port, scan.Name())
					if !slices.Contains(foundScanners, scan.Name()) {
						servicesWithScan = append(servicesWithScan,
							validScanner{
								scanner: scan,
								service: service,
								host:    host,
							})
						foundScanners = append(foundScanners, scan.Name())
					}
				}
			} else {
				out.Warn("Found service %s on port %d - found no scan", service.Name, service.Port)
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
			scanner.scanner.Run(scanner.service, scanner.host)
		}()
	}

	wg.Wait()
}
