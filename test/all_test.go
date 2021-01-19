package tusktest

//this is the test suite for tools

import (
	"testing"

	"github.com/tusklang/goat"
	"github.com/tusklang/tusk/lang/types"
)

func TestAll(t *testing.T) {
	lib, _ := goat.LoadLibrary("./tusk", types.CliParams{})

	goat.CallOatFunc(lib, "main")
}
