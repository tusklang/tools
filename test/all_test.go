package tusktest

//this is the test suite for tools

import (
	"fmt"
	"testing"

	"github.com/tusklang/goat"
	"github.com/tusklang/tusk/lang/types"
)

func TestAll(t *testing.T) {
	goat.DefOp("+", "int", "string", func(
		val1 types.TuskType,
		val2 types.TuskType,
		tuskpanic func(
			msg string,
			code int,
		) *types.TuskError,
	) (*types.TuskType, *types.TuskError) {
		fmt.Println("SSSSSSSSS")
		return nil, nil
	})

	lib, _ := goat.LoadLibrary("./tusk", types.CliParams{})

	goat.CallOatFunc(lib, "main")
}
