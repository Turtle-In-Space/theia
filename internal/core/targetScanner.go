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

	"github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// ----- Structs ----- //

// implemeted for future use
type target struct {
	name  string
	hosts []host
}

type host struct {
	hostname  string
	ipAddr    string
	services  []service
	dataDir   string
	resultDir string
}

type service struct {
	name string
	port int
}

// TODO rename, add ipAddr
type validScanner struct {
	scanner ServiceScanner
	service service
	host    host
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
	dataOutPath := filepath.Join("..", "ports.xml") // := scanTarget(ip)
	target := GetTarget(dataOutPath, targetName)
	target.addDirs()

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

func (t *target) addDirs() {
	if len(t.hosts) == 1 {
		host := &t.hosts[0]
		host.dataDir = dataDir
		host.resultDir = resultDir
	} else {
		for _, host := range t.hosts {
			host.addDirs()
		}
	}
}

// create a host and dirs for host
func (h *host) addDirs() {
	// create dirs for host
	h.dataDir = filepath.Join(dataDir, h.ipAddr)
	h.resultDir = filepath.Join(resultDir, h.ipAddr)

	helpers.CreateDir(h.dataDir)
	helpers.CreateDir(h.resultDir)
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

func queueScanners(target target) (servicesWithScan []validScanner) {
	var foundScanners []string

	// find scan for each serivce
	for _, host := range target.hosts {
		// clear scanners per host
		foundScanners = nil

		for _, service := range host.services {
			scan, ok := ScannerByServiceName(service.name)

			if ok {
				out.Info("Found service %s on port %d - using scan %s", service.name, service.port, scan.Name())
				if !slices.Contains(foundScanners, scan.Name()) {
					servicesWithScan = append(servicesWithScan,
						validScanner{
							scanner: scan,
							service: service,
							host:    host,
						})
					foundScanners = append(foundScanners, scan.Name())
				}
			} else {
				out.Warn("Found service %s on port %d - found no scan", service.name, service.port)
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
