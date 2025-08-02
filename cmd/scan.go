/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"time"

	"github.com/Turtle-In-Space/theia/internal/core"
	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/scan"
	out "github.com/Turtle-In-Space/theia/pkg/output"

	"github.com/spf13/cobra"
)

var (
	targetName string
	ipAddr     string
	startTime  time.Time
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   msg.Usage,
	Short: msg.Short,
	Long:  msg.Long,
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
	core.ScanTarget(ipAddr, targetName)
}

func startTimer() {
	startTime = time.Now()
}

func endTimer() {
	elapsed := time.Since(startTime)

	out.Success("Completed all scans in %s", elapsed.Round(time.Millisecond))
}
