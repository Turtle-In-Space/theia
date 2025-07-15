/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"os/exec"
	"path/filepath"
	"slices"

	core "github.com/Turtle-In-Space/theia/internal/core"
	"github.com/Turtle-In-Space/theia/internal/scanners"
	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/scan"
	helpers "github.com/Turtle-In-Space/theia/pkg/helpers"
	out "github.com/Turtle-In-Space/theia/pkg/output"

	"github.com/spf13/cobra"
)

var (
	targetName string
	ipAddr     string
	targetDir  string
	xmlDir     string
	resultDir  string
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   msg.Usage,
	Short: msg.Short,
	Long:  msg.Long,
	Args:  cobra.ExactArgs(1),

	// Store args, create dirs then begin scan
	Run: func(cmd *cobra.Command, args []string) {
		getArgs(args)
		initProject()
		scanTarget()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&targetName, "name", "n", "", "Name of the target")
}

// store cmd args
func getArgs(args []string) {
	ipAddr = args[0]
	targetDir = targetName
}

// begin the target scan
func scanTarget() {
	openPortScan()

	portsFile := filepath.Join(xmlDir, "ports.xml")
	services := core.GetServices(portsFile)

	scannerQueue := queueScanners(services)
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

func queueScanners(services map[int]string) (scannerQueue []scanners.ServiceScanner) {
	var foundScanners []string

	// find scan for each serivce
	for port, service := range services {
		scan, ok := scanners.ScannerBySericeName(service)

		if ok {
			out.Info("Found service %s on port %d - using scan %s", service, port, scan.Name())
			if !slices.Contains(foundScanners, scan.Name()) {
				scannerQueue = append(scannerQueue, scan)
				foundScanners = append(foundScanners, scan.Name())
			}
		} else {
			out.Warn("Found service %s on port %d - found no scan", service, port)
		}
	}

	return
}
