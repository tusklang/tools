package tusktest

//this is the test suite for tools

import (
	"os"
	"testing"

	"github.com/tusklang/goat"
	"github.com/tusklang/tusk/lang/types"
)

func TestAll(t *testing.T) {

	os.Chdir("..") //go to the parent directory (for tuskstd)

	lib, e := goat.LoadLibrary("./test/tusk", types.CliParams{})

	if e != nil {
		panic(e)
	}

	_, e2 := goat.CallOatFunc(lib, "main")

	if e2 != nil {
		(*e2).Print()
	}
}
