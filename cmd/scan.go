/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Turtle-In-Space/theia/core"
	"github.com/Turtle-In-Space/theia/output"

	"github.com/spf13/cobra"
)

// ----- Variables ----- //

var (
	scansRanCount int
	targetName    string
	startTime     time.Time
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [IP]",
	Short: "scan a target",
	Long:  "scan a target",
	Args:  cobra.ExactArgs(1),

	PreRun: func(_ *cobra.Command, _ []string) {
		startTimer()
	},

	// Store args then start scan
	Run: func(cmd *cobra.Command, args []string) {
		getArgs(args)
		startScan()
	},

	PostRun: func(_ *cobra.Command, _ []string) {
		endTimer()
	},
}

// ----- Private Functions ----- //

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&targetName, "name", "n", "", "Name of the target")
}

// store cmd args
func getArgs(args []string) {
	ipAddr = args[0]
}

// run the actual scan
func startScan() {
	if targetName == "" {
		targetName = fmt.Sprintf("theia-scan_%s", time.Now().Format(time.RFC3339))
	}

	scansRanCount = core.ScanTarget(ipAddr, targetName)
}

func startTimer() {
	startTime = time.Now()
}

func endTimer() {
	elapsed := time.Since(startTime)

	output.Success(output.Normal, "Completed %s scans in %s", scansRanCount, elapsed.Round(time.Millisecond))
}
