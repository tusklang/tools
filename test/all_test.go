package tusktest

//this is the test suite for tools

import (
	"testing"

	"github.com/tusklang/goat"
	"github.com/tusklang/tusk/lang/types"
)

func TestAll(t *testing.T) {
	lib, e := goat.LoadLibrary("./tusk", types.CliParams{})

	if e != nil {
		panic(e)
	}

	goat.CallOatFunc(lib, "main")
}
