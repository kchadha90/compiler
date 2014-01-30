// d1aefb1bfa9a074d43b1315f35c41bed
package testmylexer

import (
"cs553s2013/mylexer"
"github.com/proebsting/cs553s2013/lexer"
"testing"
)

func Test_d1aefb1bfa9a074d43b1315f35c41bed_20130120120713(t *testing.T) {
c := make(chan lexer.Token, 100)
s := `'a'
`
v := []lexer.Token{
TestToken{lexer.ERROR, "'", 0, 1, 1},
}
go mylexer.Lexer(s, c)
lexer.CheckTest(s, c, v, t)
}
