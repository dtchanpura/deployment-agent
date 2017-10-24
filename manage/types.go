package manage

import "time"

// Configuration json structure for JSON File.
type Configuration struct {
	ConfigFilePath string       `json:"config_file_path"`
	TokenSecret    string       `json:"token_secret"`
	Repositories   []Repository `json:"repositories"`
}

// Repository json structure to be stored in Configuration file.
type Repository struct {
	Name           string   `json:"name"`
	Path           string   `json:"path"`
	Token          string   `json:"token"`
	RemoteName     string   `json:"remote_name"`
	PostHookPath   string   `json:"post_hook_path"`
	LastUpdated    string   `json:"last_updated"`
	WhiteListedIPs []string `json:"whitelisted_ips"`
}

// Response defines the json structure to return as response after webhook call.
type Response struct {
	StatusCode int    `json:"status_code"`
	Ok         bool   `json:"ok"`
	Message    string `json:"message,omitempty"`
}

// TravisWebhookBody describes the payload of webhook
type TravisWebhookBody struct {
	ID            int       `json:"id"`
	Number        string    `json:"number,omitempty"`
	Type          string    `json:"type,omitempty"`
	State         string    `json:"state,omitempty"`
	Status        int       `json:"status,omitempty"` // 0 or 1 depending upon the build success or failure respectively
	Result        int       `json:"result,omitempty"` // 0 or 1 depending upon the build success or failure respectively
	StatusMessage string    `json:"status_message,omitempty"`
	ResultMessage string    `json:"result_message,omitempty"`
	StartedAt     time.Time `json:"started_at,omitempty"`
	FinishedAt    time.Time `json:"finished_at,omitempty"`
	Duration      int       `json:"duration,omitempty"`
	BuildURL      string    `json:"build_url,omitempty"`
	// CommitID       string       `json:"commit_id,omitempty"`
	Commit         string    `json:"commit,omitempty"`
	BaseCommit     string    `json:"base_commit,omitempty"`
	HeadCommit     string    `json:"head_commit,omitempty"`
	Branch         string    `json:"branch,omitempty"`
	Message        string    `json:"message,omitempty"`
	CompareURL     string    `json:"compare_url,omitempty"`
	CommittedAt    time.Time `json:"committed_at,omitempty"`
	AuthorName     string    `json:"author_name,omitempty"`
	AuthorEmail    string    `json:"author_email,omitempty"`
	CommitterName  string    `json:"committer_name,omitempty"`
	CommitterEmail string    `json:"committer_email,omitempty"`
	Tag            string    `json:"tag,omitempty"`
}
