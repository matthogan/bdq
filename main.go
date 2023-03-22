package main

import (
	"fmt"
	"os"

	"github.com/matthogan/bdq/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var blackDuckURL string
var apiToken string

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config file in the current directory and $HOME directory
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv() // Read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}
}

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
