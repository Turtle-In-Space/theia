package core

import (
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/Turtle-In-Space/theia/internal/scanners"
	helpers "github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

// TODO rename
type validScanner struct {
	scanner scanners.ServiceScanner
	port    int
}

var (
	targetDir string
	xmlDir    string
	resultDir string
	ipAddr    string
)

// begin the target scan
func ScanTarget(ip string, targetName string) {
	targetDir = targetName
	ipAddr = ip
	initProject()
	openPortScan()

	portsFile := filepath.Join(xmlDir, "ports.xml")
	services := GetServices(portsFile)

	scannerQueue := queueScanners(services)
	runScanners(scannerQueue)
}

func initProject() {
	// create dir structure
	xmlDir = filepath.Join(targetDir, "xml/")
	resultDir = filepath.Join(targetDir, "results/")

	helpers.CreateDir(targetDir)
	helpers.CreateDir(xmlDir)
	helpers.CreateDir(resultDir)

	out.Info("created dirs")
}

func openPortScan() {
	xmlOut := filepath.Join(xmlDir, "ports.xml")
	txtOut := filepath.Join(resultDir, "ports.xml")

	cmd := exec.Command("nmap", ipAddr, "-oX", xmlOut, "-oN", txtOut)
	err := cmd.Run()

	if err != nil {
		out.Error(err.Error())
	}
}

func queueScanners(services map[int]string) (servicesWithScan []validScanner) {
	var foundScanners []string

	// find scan for each serivce
	for port, serviceName := range services {
		scan, ok := scanners.ScannerByServiceName(serviceName)

		if ok {
			out.Info("Found service %s on port %d - using scan %s", serviceName, port, scan.Name())
			if !slices.Contains(foundScanners, scan.Name()) {
				servicesWithScan = append(servicesWithScan,
					validScanner{
						scanner: scan,
						port:    port,
					})
				foundScanners = append(foundScanners, scan.Name())
			}
		} else {
			out.Warn("Found service %s on port %d - found no scan", serviceName, port)
		}
	}

	return
}

func runScanners(scannerQueue []validScanner) {
	for _, scanner := range scannerQueue {
		scanner.scanner.Run(ipAddr, scanner.port)
	}
}
