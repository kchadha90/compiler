// 2e8b519f3621f75765c979941913e5b5
package testmylexer

import (
"cs553s2013/mylexer"
"github.com/proebsting/cs553s2013/lexer"
"testing"
)

func Test_2e8b519f3621f75765c979941913e5b6_20130120120713(t *testing.T) {
c := make(chan lexer.Token, 100)
s := `123.D
`
v := []lexer.Token{
TestToken{lexer.ERROR, "123.D", 0, 1, 1},
}
go mylexer.Lexer(s, c)
lexer.CheckTest(s, c, v, t)
}
