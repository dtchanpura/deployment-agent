package cmd

import (
	"github.com/dtchanpura/deployment-agent/config"
	"github.com/spf13/cobra"
)

var forceOverwrite bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration",
	Long: `This command is for initializing configuration file with minimal
config values.

Config values are mainly default and should be used only once.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)
		config.InitializeConfiguration(cfgFile, forceOverwrite)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolVarP(&forceOverwrite, "force", "f", false, "Overwrite configurations forcefully. Use with CAUTION.")
}
