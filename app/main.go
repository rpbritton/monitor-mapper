package main

import (
	"fmt"
	"os"

	"github.com/rpbritton/monitor-mapper/app/configuration"
	"github.com/sirupsen/logrus"
)

func main() {
	// parse flags
	flags, err := configuration.ParseFlags(os.Args[1:])
	if err != nil {
		configuration.PrintUsageErr(flags, err)
	}
	if flags.Help {
		configuration.PrintUsage(flags)
	}

	// set up logging
	logger := logrus.New()
	if flags.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
	logger.Debug("set up logger")

	// parse config
	config, err := configuration.ParseConfig(flags.Config)
	if err != nil {
		configuration.PrintConfigErr(err)
	}
	logger.WithField("configPath", flags.Config).Debug("parsed config file")

	fmt.Println(config)
}
