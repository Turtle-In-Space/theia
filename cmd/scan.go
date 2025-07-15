/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"github.com/Turtle-In-Space/theia/internal/core"
	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/scan"

	"github.com/spf13/cobra"
)

var (
	targetName string
	ipAddr     string
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   msg.Usage,
	Short: msg.Short,
	Long:  msg.Long,
	Args:  cobra.ExactArgs(1),

	// Store args then start scan
	Run: func(cmd *cobra.Command, args []string) {
		getArgs(args)
		startScan()
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

func startScan() {
	core.ScanTarget(ipAddr, targetName)
}
