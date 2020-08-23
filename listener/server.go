package listener

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

var (
	logger  = zerolog.New(os.Stdout)
	dateVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "project_time",
			Namespace: "agent",
			Help:      "Project Deployment time.",
		},
		[]string{
			"name",
			"uuid",
			"ip",
			"status",
		})
)

func init() {
	prometheus.MustRegister(dateVec)
}

// StartListener for starting the gin server on given host:port
func StartListener(host string, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/reload/", webHookHandler)
	mux.HandleFunc("/version", versionHandler)
	mux.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf("%s:%v", host, port)
	logger.Info().Msgf("Server started at %s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
}
