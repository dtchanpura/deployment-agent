package constants

import "time"

// SecretConstants are constants which will be used for generating random string
var SecretConstants = []string{
	"01234567",                   // include Octals (weakest)
	"89abcdef",                   // include Hexadecimals (weaker)
	"qwrtyuiopsghjklzxvnm",       // include rest lower case alphabets (weak)
	"QWERTYUIOPASDFGHJKLZXCVBNM", // ^^ include all upper case alphabets (still weak)
	"~!@#$%^&*()",                // include some symbols (seems strong)
}

// Version string to store the command version
var Version = ""

// BuildDateStr string to store the built date
var BuildDateStr = ""

// BuildDate function to format date properly
func BuildDate() string {
	if d, err := time.Parse(time.RFC3339, BuildDateStr); err == nil {
		return d.String()
	}
	return BuildDateStr
}
