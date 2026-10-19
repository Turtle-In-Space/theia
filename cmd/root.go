/*
Copyright © 2025 Elias Svensson <elias.svensson63@gmail.com>
*/
package cmd

import (
	"github.com/Turtle-In-Space/theia/output"

	"os"

	"github.com/spf13/cobra"
)

var (
	// cfgFile        string
	verbosityCount int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "theia",
	Short: "Automated host scanner",
	Long: `
 ______  __ __    ___  ____   ____ 
|      ||  |  |  /  _]|    | /    |
|      ||  |  | /  [_  |  | |  o  |
|_|  |_||  _  ||    _] |  | |     |
  |  |  |  |  ||   [_  |  | |  _  |
  |  |  |  |  ||     | |  | |  |  |
  |__|  |__|__||_____||____||__|__|
                                   
An automated host scanner. It will scan given targets and then enumerate discovered services.`,

	Version: "v0.2.0",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		output.SetThreshold(verbosityCount)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().CountVarP(&verbosityCount, "verbose", "v", "Enable verbose output. Repeat for more verbosity. (max -vv)")

	// cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.theia.yaml)")
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)
//
// 		// Search config in home directory with name ".theia" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".theia")
// 	}
//
// 	viper.AutomaticEnv() // read in environment variables that match
//
// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }
