/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"github.com/Turtle-In-Space/theia/core"
	"github.com/spf13/cobra"
)

var ()

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [IP] [PATH]",
	Short: "create a .env file for the target",
	Long:  "create a .env file for the target",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		core.CreateEnvFile(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
