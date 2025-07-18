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

	helpers "github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// ----- Structs ----- //

// implemeted for future use
type target struct {
	name  string
	hosts []host
}

type host struct {
	hostname     string
	ipAddr       string
	services     []service
	dataFolder   string
	resultFolder string
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

// ----- Functions ----- //

// begin the target scan
func ScanTarget(ip string, targetName string) {
	target := createTarget(ip, targetName)
	target.createTargetStructure()

	scannerQueue := queueScanners(target)
	runScanners(scannerQueue)
}

func createTarget(ip, name string) target {
	dataOutPath := filepath.Join("nmap.xml") // := scanTarget(ip)
	return GetTarget(dataOutPath, name)
}

func (t *target) createTargetStructure() {
	helpers.CreateDir(t.name)
	os.Chdir(t.name)

	// create dir structure
	dataDir = filepath.Clean("data/")
	resultDir = filepath.Clean("results/")

	helpers.CreateDir(dataDir)
	helpers.CreateDir(resultDir)

	out.Info("created dirs")

	for _, host := range t.hosts {
		host.addDirs()
	}
}

// create a host and dirs for host
func (h *host) addDirs() {
	// create dirs for host
	h.dataFolder = filepath.Join(dataDir, h.ipAddr)
	h.resultFolder = filepath.Join(resultDir, h.ipAddr)

	helpers.CreateDir(h.dataFolder)
	helpers.CreateDir(h.resultFolder)
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
