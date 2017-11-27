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
	Name       string        `yaml:"name"`
	UUID       string        `yaml:"uuid"`
	WorkDir    string        `yaml:"work_dir"`
	RemotePath string        `yaml:"remote_path"` // For downloading
	PostHook   string        `yaml:"post_hook,omitempty"`
	PreHook    string        `yaml:"pre_hook,omitempty"`
	ErrorHook  string        `yaml:"error_hook,omitempty"`
	Secret     string        `yaml:"secret"`
	Tokens     []TokenDetail `yaml:"tokens"`
}

// TokenDetail is for allowing multiple ips to access same
// repository with different tokens
type TokenDetail struct {
	Token              string `yaml:"token"`
	WhitelistedNetwork string `yaml:"whitelist_net"` // CIDR notation
}
