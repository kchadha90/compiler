// 021399470f960dc3f467b49dcaf3d615
package testmylexer

import (
	"cs553s2013/compiler"
	"testing"
	//"fmt"
)

func Test_simple_error(t *testing.T) {
	oberon := `
MODULE mymod;
	VAR 
		i : INTEGER;
		b : LOL;	
	
END mymod.
`
	expectedStatus := true
	ok, msg := compiler.Checker(oberon)
	if ok != expectedStatus {
		t.Errorf("FAIL:  expected %v status.  Msg: %v\n", expectedStatus, msg)
	}
}

func Test_simple_ok(t *testing.T) {
	oberon := `
MODULE mymod
END mymod.
`
	expectedStatus := true
	ok, msg := compiler.Checker(oberon)
	if ok != expectedStatus {
		t.Errorf("FAIL:  expected %v status.  Msg: %v\n", expectedStatus, msg)
	}
}
