package config

// Configuration for whole configuration file
type Configuration struct {
	ServeConfig    Serve     `yaml:"serve" json:"serve"`
	ProjectConfigs []Project `yaml:"projects,omitempty" json:"projects,omitempty"`
}

// Serve for server (listener) related configurations
type Serve struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}

// Project for all projects configurations
type Project struct {
	Name          string        `yaml:"name" json:"name"`
	UUID          string        `yaml:"uuid" json:"uuid"`
	WorkDir       string        `yaml:"work_dir" json:"work_dir"`
	PostHook      string        `yaml:"post_hook,omitempty" json:"post_hook,omitempty"`
	PostHookArgs  []string      `yaml:"post_hook_args,omitempty" json:"post_hook_args,omitempty"`
	PreHook       string        `yaml:"pre_hook,omitempty" json:"pre_hook,omitempty"`
	PreHookArgs   []string      `yaml:"pre_hook_args,omitempty" json:"pre_hook_args,omitempty"`
	ErrorHook     string        `yaml:"error_hook,omitempty" json:"error_hook,omitempty"`
	ErrorHookArgs []string      `yaml:"error_hook_args,omitempty" json:"error_hook_args,omitempty"`
	Secret        string        `yaml:"secret" json:"secret"`
	Tokens        []TokenDetail `yaml:"tokens" json:"tokens"`
	Hooks         []Hook        `yaml:"hooks,omitempty" json:"hooks,omitempty"`
	// Following part has been removed as we will be adding all related things in PreHook or PostHook
	// RemotePath    string        `yaml:"remote_path" json:"remote_path"` // For downloading
}

// Hook is a new struct for replacing PreHook and PostHook
type Hook struct {
	FilePath string `yaml:"file_path" json:"file_path"`                   // path to file to be executed
	MaxArgs  int    `yaml:"max_args,omitempty" json:"max_args,omitempty"` // For limiting number of arguments
	// -1 allows all arguments
	// TODO: Hook specific WorkDir to be added
}

// TokenDetail is for allowing multiple ips to access same
// repository with different tokens
type TokenDetail struct {
	Token              string `yaml:"token" json:"token"`
	WhitelistedNetwork string `yaml:"whitelistnet" json:"whitelistnet"` // CIDR notation
	Name               string `yaml:"name" json:"name"`
}
