/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/scan"
	helpers "github.com/Turtle-In-Space/theia/pkg"

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

	Run: func(cmd *cobra.Command, args []string) {
		ipAddr = args[0]
		targetDir = targetName

		initProject()
		scanTarget()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&targetName, "name", "n", "", "Name of the target")
}

func scanTarget() {
	openPortScan()

}

func initProject() {
	// create dir structure
	xmlDir = fmt.Sprintf("%s/xml", targetDir)
	resultDir = fmt.Sprintf("%s/results", targetDir)

	helpers.CreateDir(targetDir)
	helpers.CreateDir(xmlDir)
	helpers.CreateDir(resultDir)

	fmt.Println("[*] made dirs")
}

func openPortScan() {

	outputPath := fmt.Sprintf("%s/ports.xml", xmlDir)

	cmd := exec.Command("nmap", ipAddr, "-oX", outputPath)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
