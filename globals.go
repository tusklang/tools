package tools

import (
	"flag"
	"fmt"
)

const (
	//TuskMajor is the major version of the Tusk suite
	TuskMajor = 0
	//TuskMinor is the minor version of the Tusk suite
	TuskMinor = 0
	//TuskBug is the bug version of the Tusk suite
	TuskBug = 0
)

//Usagef is the default Usage message in the golang "flag" package, but instead it prints the Tusk addon name, instead of the executable
func Usagef(addon string) func() {
	//set a custom "Usage of ..." message
	return func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s\n", addon)
		flag.PrintDefaults()
	}
}
