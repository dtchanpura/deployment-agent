package listener

// Response defines the json structure to return as response after webhook call.
type Response struct {
	StatusCode int         `json:"status_code"`
	Ok         bool        `json:"ok"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Version    string      `json:"version,omitempty"`
	BuildDate  string      `json:"build_date,omitempty"`
}
