package lexer

import (
	"testing"
)

func CheckTest(s string, c chan Token, v []Token, t *testing.T) {
	var i int
	var x Token

	for i, x = 0, <-c; i < len(v); i, x = i+1, <-c {
		y := v[i]
		if x.Enum() != y.Enum() || x.Line() != y.Line() || x.Column() != y.Column() || (x.Enum() != EOF && x.Enum() != ERROR && x.Value() != y.Value()) {
			t.Errorf("mismatch: Found(%s,%s,%d,%d) Expected(%s,%s,%d,%d)", 
			Enum2String(x.Enum()), x.Value(), x.Line(), x.Column(),
			Enum2String(y.Enum()), y.Value(), y.Line(), y.Column())
			return
		}
		if x.Enum() == EOF || x.Enum() == ERROR {
			return
		}
	}

	t.Fatal("Email Todd Proebsting---this shouldn't happend")
}
