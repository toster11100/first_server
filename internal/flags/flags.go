package flags

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	Port *string
}

func New() Flags {
	flags := Flags{}

	flags.Port = pflag.StringP("port", "p", ":8080", "port")
	pflag.Parse()

	return flags
}
