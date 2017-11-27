// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	defaultHost string
	defaultPort int
	host        string
	port        int
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve command for starting listener.",
	Long:  `Serve command is to start the listener to access API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)

		// If host or port is default use the one from configuration
		if host == defaultHost {
			host = viper.GetString("serve.host")
		}
		if port == defaultPort {
			port = viper.GetInt("serve.port")
		}

		fmt.Println("serve called")
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	defaultHost = ""
	defaultPort = 8000
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringVarP(&host, "host", "H", defaultHost, "Host where to listen")
	serveCmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Port where to listen")
}
