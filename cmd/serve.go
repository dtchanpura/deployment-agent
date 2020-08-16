package cmd

import (
	"github.com/dtchanpura/deployment-agent/listener"
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

		// fmt.Println("serve called")
		listener.StartListener(host, port)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	defaultHost = ""
	defaultPort = 8000

	serveCmd.Flags().StringVarP(&host, "host", "H", defaultHost, "Host where to listen")
	serveCmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Port where to listen")
}
