package suite

import (
	"flag"
	"fmt"
)

const (
	//OmmSuiteMajor is the major version of the Omm suite
	OmmSuiteMajor = 0
	//OmmSuiteMinor is the minor version of the Omm suite
	OmmSuiteMinor = 0
	//OmmSuiteBug is the bug version of the Omm suite
	OmmSuiteBug = 0
)

//Usagef is the default Usage message in the golang "flag" package, but instead it prints the Omm addon name, instead of the executable
func Usagef(addon string) func() {
	//set a custom "Usage of ..." message
	return func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s\n", addon)
		flag.PrintDefaults()
	}
}
