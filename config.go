package eguid

import (
	"github.com/speps/go-hashids/v2"
)

type config struct {
	Salt     string
	Length   int
	Alphabet string
}

// DefaultConfig ...
func DefaultConfig() *config {
	return &config{
		Salt:     "",
		Length:   16,
		Alphabet: hashids.DefaultAlphabet,
	}
}
