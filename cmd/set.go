/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"github.com/Turtle-In-Space/theia/internal/core"
	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/set"
	"github.com/spf13/cobra"
)

var ()

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   msg.Usage,
	Short: msg.Short,
	Long:  msg.Long,
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		core.CreateEnvFile(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
