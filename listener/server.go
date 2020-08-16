package listener

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger = zerolog.New(os.Stdout)
)

// StartListener for starting the gin server on given host:port
func StartListener(host string, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/reload/", webHookHandler)
	mux.HandleFunc("/version", versionHandler)
	addr := fmt.Sprintf("%s:%v", host, port)
	logger.Info().Msgf("Server started at %s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
}
