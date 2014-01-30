// 3aa87855bc07809d26729c80731b2bf5
package testmylexer

import (
"cs553s2013/mylexer"
"github.com/proebsting/cs553s2013/lexer"
"testing"
)

func Test_3aa87855bc07809d26729c80731b2bf5_20130120120713(t *testing.T) {
c := make(chan lexer.Token, 100)
s := `"ab
`
v := []lexer.Token{
TestToken{lexer.ERROR, "string error", 0, 1, 1},
}
go mylexer.Lexer(s, c)
lexer.CheckTest(s, c, v, t)
}
