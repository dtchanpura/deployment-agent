// Copyright © 2017 Darshil Chanpura <dtchanpura@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
