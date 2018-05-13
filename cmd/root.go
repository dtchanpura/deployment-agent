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
	"fmt"
	"os"
	"path"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/fsnotify/fsnotify"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var watchEnabled bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "deployment-agent",
	Short: "Deployment agent for simple deployments",
	Long: `Deployment agent for initializing, adding and modifying configurations.

deployment-agent is a cli tool to manage paths, hooks and tokens for accessing
deployment resources which can be useful to be called from Jenkins or other CI tools.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deployment-agent.yaml)")
	RootCmd.PersistentFlags().BoolVar(&watchEnabled, "watch-config", false, "Watch config file for changes")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("verbose", "v", false, "Enables debugging")
	RootCmd.Flags().BoolP("version", "V", false, "Displays version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".deployment-agent" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".deployment-agent")
		// viper.SetConfigType("yaml")
		viper.SetDefault("serve.host", "")
		viper.SetDefault("serve.port", 8000)
		cfgFile = path.Join(home, ".deployment-agent.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		viper.UnmarshalKey("serve", &config.StoredServe)
		// err := viper.UnmarshalKey("projects", &config.StoredProjects)
		config.DecodeProjectConfiguration(viper.AllSettings())
		// err = viper.Unmarshal(&config.StoredConfiguration)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(config.StoredConfiguration)
		// viper.Unmarshal(&config.StoredConfiguration)
		// config.ViperConfiguration = viper.Sub("projects")
	}
	if watchEnabled {
		// Watch part
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			viper.UnmarshalKey("serve", &config.StoredServe)
			// viper.UnmarshalKey("projects", &config.StoredProjects)
			config.DecodeProjectConfiguration(viper.AllSettings())
		})

	}
}
