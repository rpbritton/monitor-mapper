package configuration

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

var ErrUnexpectedArgs = errors.New("unexpected args")

type Flags struct {
	Config     string
	Debug      bool
	List       bool
	Background bool
	Help       bool

	set *flag.FlagSet
}

func ParseFlags(args []string) (Flags, error) {
	var flags Flags

	// create set
	flags.set = flag.NewFlagSet("monitormapper", flag.ContinueOnError)

	// define
	flags.set.StringVar(&flags.Config, "config", "~/.config/monitor-mapper/config.yaml", "path to the config file")
	flags.set.BoolVar(&flags.Debug, "debug", false, "enable debug logging")
	flags.set.BoolVar(&flags.Debug, "list", false, "list all connected monitors")
	flags.set.BoolVar(&flags.Debug, "background", false, "start background mode")

	// parse
	flags.set.SetOutput(io.Discard)
	err := flags.set.Parse(args)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			flags.Help = true
			return flags, nil
		} else {
			return flags, fmt.Errorf("parse error: %w", err)
		}
	}

	// validate
	if flags.set.NArg() > 0 {
		return flags, fmt.Errorf("%w: %v", ErrUnexpectedArgs, flags.set.Args())
	}

	return flags, nil
}
