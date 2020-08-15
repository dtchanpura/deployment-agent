package listener

import (
	"fmt"
	"log"
	"net/http"
)

// StartListener for starting the gin server on given host:port
func StartListener(host string, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/reload/", webHookHandler)
	mux.HandleFunc("/version", versionHandler)
	addr := fmt.Sprintf("%s:%v", host, port)
	log.Println("Server started at", addr)
	log.Fatalln(http.ListenAndServe(addr, mux))
}
