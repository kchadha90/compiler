// de5303bd2e4cd113654e807f6361d0d0
package testmylexer

import (
"cs553s2013/mylexer"
"github.com/proebsting/cs553s2013/lexer"
"testing"
)

func Test_de5303bd2e4cd113654e807f6361d0d0_20130122210446(t *testing.T) {
c := make(chan lexer.Token, 100)
s := `MODULE Sample1;

  IMPORT c := Console;

BEGIN
  c.String("Hello, World!"); c.Ln;
END Sample1.`

v := []lexer.Token{
TestToken{lexer.MODULE, "MODULE", 0, 1, 1},
TestToken{lexer.IDENT, "Sample1", 0, 1, 8},
TestToken{lexer.SEMICOLON, ";", 0, 1, 15},
TestToken{lexer.IMPORT, "IMPORT", 0, 3, 3},
TestToken{lexer.IDENT, "c", 0, 3, 10},
TestToken{lexer.COLONEQUAL, ":=", 0, 3, 12},
TestToken{lexer.IDENT, "Console", 0, 3, 15},
TestToken{lexer.SEMICOLON, ";", 0, 3, 22},
TestToken{lexer.BEGIN, "BEGIN", 0, 5, 1},
TestToken{lexer.IDENT, "c", 0, 6, 3},
TestToken{lexer.DOT, ".", 0, 6, 4},
TestToken{lexer.IDENT, "String", 0, 6, 5},
TestToken{lexer.LPAREN, "(", 0, 6, 11},
TestToken{lexer.STRING, "\"Hello, World!\"", 0, 6, 12},
TestToken{lexer.RPAREN, ")", 0, 6, 27},
TestToken{lexer.SEMICOLON, ";", 0, 6, 28},
TestToken{lexer.IDENT, "c", 0, 6, 30},
TestToken{lexer.DOT, ".", 0, 6, 31},
TestToken{lexer.IDENT, "Ln", 0, 6, 32},
TestToken{lexer.SEMICOLON, ";", 0, 6, 34},
TestToken{lexer.END, "END", 0, 7, 1},
TestToken{lexer.IDENT, "Sample1", 0, 7, 5},
TestToken{lexer.DOT, ".", 0, 7, 12},
TestToken{lexer.EOF, "<EOF>", 0, 7, 13},
}
go mylexer.Lexer(s, c)
lexer.CheckTest(s, c, v, t)
}
