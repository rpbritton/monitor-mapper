package configuration

import (
	"fmt"
	"os"
)

var usageOutput = os.Stderr

func PrintUsage(flags Flags) {
	PrintUsageErr(flags, nil)
}

func PrintUsageErr(flags Flags, err error) {
	// check error
	if err != nil {
		fmt.Fprintf(usageOutput, "Failed to parse arguments: %v\n\n", err)
	}

	// print usage
	fmt.Fprintf(usageOutput, "Usage of monitormapper:\n")
	flags.set.SetOutput(usageOutput)
	flags.set.PrintDefaults()

	// exit
	Exit(err)
}

func PrintConfigErr(err error) {
	// print error
	fmt.Fprintf(usageOutput, "Bad config file: %v\n\n", err)

	// exit
	Exit(err)
}

func Exit(err error) {
	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
