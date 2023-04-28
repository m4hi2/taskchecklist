/*
Copyright © 2023 Mahir Labib Chowdhury dev.mahirchy@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskchecklist",
	Short: "Server application for a checklist application",
	Long: `Server application for a checklist application.

The primary goal of this application is to have a checklist
for repeated tasks. It is to help with not missing any
important steps of the tasks.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName(".taskchecklist")

	env := "local"

	if envVar := os.Getenv("SERVICE_ENV"); envVar != "" {
		env = strings.ToLower(envVar)
	}
	viper.Set("server.env", env)

	if env == "local" {
		viper.AddConfigPath("./config")
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".taskchecklist" (without extension).
		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		cobra.CheckErr(err)
	}
}
