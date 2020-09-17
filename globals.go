package kore

import (
	"flag"
	"fmt"
)

const (
	//KoreMajor is the major version of the Tusk suite
	KoreMajor = 0
	//KoreMinor is the minor version of the Tusk suite
	KoreMinor = 0
	//KoreBug is the bug version of the Tusk suite
	KoreBug = 0
)

//Usagef is the default Usage message in the golang "flag" package, but instead it prints the Tusk addon name, instead of the executable
func Usagef(addon string) func() {
	//set a custom "Usage of ..." message
	return func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s\n", addon)
		flag.PrintDefaults()
	}
}
