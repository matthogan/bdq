package cmd

import (
	"fmt"
	"strings"

	"github.com/blackducksoftware/hub-client-go/hubapi"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use the Black Duck client to perform operations, e.g., get all projects
		projects, err := blackDuckClient.ListProjects(&hubapi.GetListOptions{})
		if err != nil {
			fmt.Printf("Error listing projects: %v\n", err)
			return
		}

		// Print the projects
		fmt.Printf("Found %d projects:\n", projects.TotalCount)
		for _, project := range projects.Items {
			projectID := extractProjectID(project.Meta.Href)
			fmt.Printf("Project: %s (Href: %s) (ID: %s)\n", project.Name, project.Meta.Href, projectID)
		}
	},
}

func init() {
	RootCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func extractProjectID(href string) string {
	// This assumes that the project ID is the last segment of the URL
	segments := strings.Split(href, "/")
	return segments[len(segments)-1]
}
