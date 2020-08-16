package cmd

import (
	"fmt"

	"github.com/dtchanpura/deployment-agent/constants"
	"github.com/spf13/cobra"
)

var (
	version   string
	buildDate string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version",
	Long:  `Displays the version of current build.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Date: %s\n", buildDate)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	version = constants.Version
	buildDate = constants.BuildDate()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
