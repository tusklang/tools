package tusktest

//this is the test suite for tools

import (
	"fmt"
	"testing"

	"github.com/tusklang/goat"
	"github.com/tusklang/tusk/lang/types"
)

func TestAll(t *testing.T) {

	goat.DefNative("testing", [][]string{
		[]string{
			"string",
		},
	}, func(args []*types.TuskType, tuskpanic goat.Panic) (*types.TuskType, *types.TuskError) {
		fmt.Println("TESTINGGGGGGGGGGGGG")

		return nil, nil
	})

	lib, _ := goat.LoadLibrary("./tusk", types.CliParams{})

	goat.CallOatFunc(lib, "main")
}
