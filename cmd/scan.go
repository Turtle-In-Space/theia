/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Turtle-In-Space/theia/internal/core"
	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/scan"
	out "github.com/Turtle-In-Space/theia/pkg/output"

	"github.com/spf13/cobra"
)

var (
	startTime time.Time
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

	core.ScanTarget(ipAddr, targetName)
}

func startTimer() {
	startTime = time.Now()
}

func endTimer() {
	elapsed := time.Since(startTime)

	out.Success("Completed all scans in %s", elapsed.Round(time.Millisecond))
}
