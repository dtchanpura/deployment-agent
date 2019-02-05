package constants

const (
	// ErrorFileExists if file already exists and overwrite flag is off.
	ErrorFileExists = "file already exists"
	// ErrorFileNotExecutable if file is not executable.
	ErrorFileNotExecutable = "file is not executable"
	// ErrorNoProjectFound if project is not found.
	ErrorNoProjectFound = "project not found"
	// ErrorProjectAlreadyExists Error if project with same UUID already exists
	ErrorProjectAlreadyExists = "name already exists" // RARE
	// ErrorInvalidConfiguration if project configuration is invalid
	ErrorInvalidConfiguration = "project configuration is invalid please check the configuration file."
	// ErrorInvalidUUID if invalid uuid is used
	ErrorInvalidUUID = "invalid uuid"
)
