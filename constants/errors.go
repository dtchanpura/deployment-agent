package constants

const (
	// ErrorFileExists if file already exists and overwrite flag is off.
	ErrorFileExists = "File already exists."
	// ErrorFileNotExecutable if file is not executable.
	ErrorFileNotExecutable = "File is not executable."
	// ErrorNoProjectFound if project is not found.
	ErrorNoProjectFound = "Project not found"
	// ErrorProjectAlreadyExists Error if project with same UUID already exists
	ErrorProjectAlreadyExists = "Project with same UUID/Name already exists" // RARE
	// ErrorInvalidConfiguration if project configuration is invalid
	ErrorInvalidConfiguration = "Project configuration is invalid please check the configuration file."
)
