/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	msg "github.com/Turtle-In-Space/theia/internal/text/cmd/set"
	"github.com/spf13/cobra"
)

var ()

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   msg.Usage,
	Short: msg.Short,
	Long:  msg.Long,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		CreateEnvFile()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

}

func CreateEnvFile() {

}
