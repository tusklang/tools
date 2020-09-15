package ommtest

//this is the test suite for omm

import (
	"testing"

	_ "goat"
	_ "omm/lang/compiler"
	_ "omm/lang/interpreter"
	_ "omm/lang/types"
	_ "omm/native"
	_ "undra/server"
)

func TestAll(t *testing.T) {

}
