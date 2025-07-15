package core

import (
	"os"
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/Turtle-In-Space/theia/internal/scanners"
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
	ipAddr   string
	services []service
}

type service struct {
	name string
	port int
}

// TODO rename, add ipAddr
type validScanner struct {
	scanner scanners.ServiceScanner
	ipAddr  string
	port    int
}

// ----- Variables ----- //

var (
	xmlDir    string
	resultDir string
)

// ----- Functions ----- //

// begin the target scan
func ScanTarget(ip string, targetName string) {
	target := createTarget(targetName)
	initProject(target)
	target.addHosts(ip)

	scannerQueue := queueScanners(target)
	runScanners(scannerQueue)
}

func createTarget(name string) target {

	return target{
		name:  name,
		hosts: nil,
	}
}

// add all hosts to target
func (t *target) addHosts(ip string) {
	t.hosts = []host{createHost(ip)}
}

// create a host and dirs for host
func createHost(ip string) host {
	// create dirs for host
	xml := filepath.Join(xmlDir, ip)
	results := filepath.Join(resultDir, ip)

	helpers.CreateDir(xml)
	helpers.CreateDir(results)

	return host{
		ipAddr:   ip,
		services: scanAllPorts(ip),
	}
}

func initProject(target target) {
	helpers.CreateDir(target.name)
	os.Chdir(target.name)

	// create dir structure
	xmlDir = filepath.Clean("xml/")
	resultDir = filepath.Clean("results/")

	helpers.CreateDir(xmlDir)
	helpers.CreateDir(resultDir)

	out.Info("created dirs")
}

func scanAllPorts(ip string) []service {
	xmlOut := filepath.Join(xmlDir, "ports.xml")
	txtOut := filepath.Join(resultDir, "ports.txt")

	cmd := exec.Command("nmap", ip, "-oX", xmlOut, "-oN", txtOut)
	err := cmd.Run()

	if err != nil {
		out.Error(err.Error())
	}

	portsFile := filepath.Join(xmlDir, "ports.xml")
	return GetServices(portsFile)
}

func queueScanners(target target) (servicesWithScan []validScanner) {
	var foundScanners []string

	// find scan for each serivce
	for _, host := range target.hosts {
		for _, service := range host.services {
			scan, ok := scanners.ScannerByServiceName(service.name)

			if ok {
				out.Info("Found service %s on port %d - using scan %s", service.name, service.port, scan.Name())
				if !slices.Contains(foundScanners, scan.Name()) {
					servicesWithScan = append(servicesWithScan,
						validScanner{
							scanner: scan,
							ipAddr:  host.ipAddr,
							port:    service.port,
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

func runScanners(scannerQueue []validScanner) {
	for _, scanner := range scannerQueue {
		scanner.scanner.Run(scanner.ipAddr, scanner.port)
	}
}
