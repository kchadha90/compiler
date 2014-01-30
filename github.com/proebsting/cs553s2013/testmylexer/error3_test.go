// 808479e2fef3695589dc6d22d906aac9
package testmylexer

import (
"cs553s2013/mylexer"
"github.com/proebsting/cs553s2013/lexer"
"testing"
)

func Test_808479e2fef3695589dc6d22d906aac9_20130120120713(t *testing.T) {
c := make(chan lexer.Token, 100)
s := `123.E+
`
v := []lexer.Token{
TestToken{lexer.ERROR, "123.E+", 0, 1, 1},
}
go mylexer.Lexer(s, c)
lexer.CheckTest(s, c, v, t)
}
