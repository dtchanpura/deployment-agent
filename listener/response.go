package listener

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response defines the json structure to return as response after webhook call.
type Response struct {
	StatusCode int         `json:"status_code"`
	Ok         bool        `json:"ok"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Version    string      `json:"version,omitempty"`
	BuildDate  string      `json:"build_date,omitempty"`
}

func (r *Response) write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if r.StatusCode >= 100 && r.StatusCode < 600 {
		w.WriteHeader(r.StatusCode)
	}
	e := json.NewEncoder(w)
	if err := e.Encode(r); err != nil {
		fmt.Fprintf(w, `{"message":"%s", "ok":false"}`, err.Error())
	}
}
