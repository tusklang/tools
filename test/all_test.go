package katest

//this is the test suite for tools

import (
	"testing"

	_ "github.com/tusklang/goat"
	_ "github.com/tusklang/tusk/lang/compiler"
	_ "github.com/tusklang/tusk/lang/interpreter"
	_ "github.com/tusklang/tusk/lang/types"
	_ "github.com/tusklang/undra/server"
)

func TestAll(t *testing.T) {

}
