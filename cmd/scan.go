/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	core "github.com/Turtle-In-Space/theia/internal/core"
	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/scan"
	helpers "github.com/Turtle-In-Space/theia/pkg/helpers"

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

	for key, val := range services {
		fmt.Printf("[*] Found %s running on port %d\n", val, key)
	}
}

func initProject() {
	// create dir structure
	xmlDir = filepath.Join(targetDir, "xml/")
	resultDir = filepath.Join(targetDir, "results/")

	helpers.CreateDir(targetDir)
	helpers.CreateDir(xmlDir)
	helpers.CreateDir(resultDir)

	fmt.Println("[*] made dirs")
}

func openPortScan() {
	xmlOut := filepath.Join(xmlDir, "ports.xml")
	txtOut := filepath.Join(resultDir, "ports.xml")

	cmd := exec.Command("nmap", ipAddr, "-oX", xmlOut, "-oN", txtOut)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
