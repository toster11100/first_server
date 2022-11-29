package flags

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	flagset pflag.FlagSet
	Port    *string
}

func New() Flags {
	flags := Flags{}

	flags.flagset = *pflag.NewFlagSet("main", pflag.ExitOnError)
	flags.Port = flags.flagset.StringP("port", "p", ":8080", "port")

	return flags
}

func (f *Flags) Parse(args []string) error {
	return f.flagset.Parse(args)
}
