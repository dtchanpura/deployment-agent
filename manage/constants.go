package manage

// TokenConstants are constants which will be used for generating random string
var TokenConstants = []string{
	"01234567",                   // include Octals (weakest)
	"89abcdef",                   // include Hexadecimals (weaker)
	"qwrtyuiopsghjklzxvnm",       // include rest lower case alphabets (weak)
	"QWERTYUIOPASDFGHJKLZXCVBNM", // ^^ include all upper case alphabets (still weak)
	"~!@#$%^&*()",                // include some symbols (seems strong)
}
