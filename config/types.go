package config

// Configuration for whole configuration file
type Configuration struct {
	ServeConfig    Serve     `yaml:"serve"`
	ProjectConfigs []Project `yaml:"projects,omitempty"`
}

// Serve for server (listener) related configurations
type Serve struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Project for all projects configurations
type Project struct {
	Path       string `yaml:"path"`
	RemotePath string `yaml:"remote_path"`
	PostHook   string `yaml:"post_hook"`
	PreHook    string `yaml:"pre_hook"`
	ErrorHook  string `yaml:"error_hook"`
}
