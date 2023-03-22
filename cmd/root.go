package cmd

import (
	"fmt"
	"time"

	"github.com/blackducksoftware/hub-client-go/hubclient"
	"github.com/spf13/cobra"
)

var cfgFile string
var blackDuckURL string
var apiToken string
var blackDuckClient *hubclient.Client

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	RootCmd.PersistentFlags().StringVarP(&blackDuckURL, "url", "u", "", "https://... black duck endpoint")
	RootCmd.PersistentFlags().StringVarP(&apiToken, "token", "t", "", "api token from the black duck interface")
}

var RootCmd = &cobra.Command{
	Use:   "bdq -u https://my.blackduck.url -t ...",
	Short: "Black Duck Client",
	Long:  `Placeholder for commonly used functionality`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { // always auth
		fmt.Println("authenticating...")
		// Authenticate and create a new Black Duck client
		blackDuckClient, err := newBlackDuckClient()
		if err != nil {
			fmt.Printf("Error creating Black Duck client: %v\n", err)
			return err
		}
		err, status := blackDuckClient.CheckHubLiveness()
		if err != nil {
			fmt.Printf("Liveness check failed: %v\n", err)
			return err
		}
		fmt.Printf("Live: %v ...\n", status.Healthy)
		err, status = blackDuckClient.CheckHubReadiness()
		if err != nil {
			fmt.Printf("Readiness check failed: %v\n", err)
			return err
		}
		fmt.Printf("Ready: %v\n !", status.Healthy)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ver, err := blackDuckClient.CurrentVersion()
		if err != nil {
			fmt.Printf("Version read failed: %v\n", err)
			return
		}
		fmt.Printf("Version: %v\n", ver.Version)
	},
}

func newBlackDuckClient() (*hubclient.Client, error) {
	// Create a new Black Duck client using the NewWithApiToken function
	debugFlags := hubclient.HubClientDebug(1)
	timeout := 5 * time.Second

	blackDuckClient, err := hubclient.NewWithApiToken(blackDuckURL, apiToken, debugFlags, timeout)
	if err != nil {
		return nil, err
	}

	return blackDuckClient, nil
}
