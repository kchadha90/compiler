// b40ed0109d32f2a975a7f5260ca12947
package testmylexer

import (
"cs553s2013/mylexer"
"github.com/proebsting/cs553s2013/lexer"
"testing"
)

func Test_b40ed0109d32f2a975a7f5260ca12947_20130122210446(t *testing.T) {
c := make(chan lexer.Token, 100)
s := `MODULE Sample2;

  IMPORT Memory, c := Console;

  VAR arr: ARRAY 20 OF CHAR;
    fixarr: POINTER TO ARRAY 20 OF CHAR;
    dynarr: POINTER TO ARRAY OF CHAR;

BEGIN
  arr := "Hello, World!";
  c.String(arr); c.Ln;

  NEW(fixarr);
  fixarr^ := "Hello, World!";
  c.String(fixarr^); c.Ln;
 
  NEW(dynarr, 20);
  dynarr^ := "Hello, World!";
  c.String(dynarr^); c.Ln;
END Sample2.`
v := []lexer.Token{
TestToken{lexer.MODULE, "MODULE", 0, 1, 1},
TestToken{lexer.IDENT, "Sample2", 0, 1, 8},
TestToken{lexer.SEMICOLON, ";", 0, 1, 15},
TestToken{lexer.IMPORT, "IMPORT", 0, 3, 3},
TestToken{lexer.IDENT, "Memory", 0, 3, 10},
TestToken{lexer.COMMA, ",", 0, 3, 16},
TestToken{lexer.IDENT, "c", 0, 3, 18},
TestToken{lexer.COLONEQUAL, ":=", 0, 3, 20},
TestToken{lexer.IDENT, "Console", 0, 3, 23},
TestToken{lexer.SEMICOLON, ";", 0, 3, 30},
TestToken{lexer.VAR, "VAR", 0, 5, 3},
TestToken{lexer.IDENT, "arr", 0, 5, 7},
TestToken{lexer.COLON, ":", 0, 5, 10},
TestToken{lexer.ARRAY, "ARRAY", 0, 5, 12},
TestToken{lexer.INTEGER, "20", 0, 5, 18},
TestToken{lexer.OF, "OF", 0, 5, 21},
TestToken{lexer.IDENT, "CHAR", 0, 5, 24},
TestToken{lexer.SEMICOLON, ";", 0, 5, 28},
TestToken{lexer.IDENT, "fixarr", 0, 6, 5},
TestToken{lexer.COLON, ":", 0, 6, 11},
TestToken{lexer.POINTER, "POINTER", 0, 6, 13},
TestToken{lexer.TO, "TO", 0, 6, 21},
TestToken{lexer.ARRAY, "ARRAY", 0, 6, 24},
TestToken{lexer.INTEGER, "20", 0, 6, 30},
TestToken{lexer.OF, "OF", 0, 6, 33},
TestToken{lexer.IDENT, "CHAR", 0, 6, 36},
TestToken{lexer.SEMICOLON, ";", 0, 6, 40},
TestToken{lexer.IDENT, "dynarr", 0, 7, 5},
TestToken{lexer.COLON, ":", 0, 7, 11},
TestToken{lexer.POINTER, "POINTER", 0, 7, 13},
TestToken{lexer.TO, "TO", 0, 7, 21},
TestToken{lexer.ARRAY, "ARRAY", 0, 7, 24},
TestToken{lexer.OF, "OF", 0, 7, 30},
TestToken{lexer.IDENT, "CHAR", 0, 7, 33},
TestToken{lexer.SEMICOLON, ";", 0, 7, 37},
TestToken{lexer.BEGIN, "BEGIN", 0, 9, 1},
TestToken{lexer.IDENT, "arr", 0, 10, 3},
TestToken{lexer.COLONEQUAL, ":=", 0, 10, 7},
TestToken{lexer.STRING, "\"Hello, World!\"", 0, 10, 10},
TestToken{lexer.SEMICOLON, ";", 0, 10, 25},
TestToken{lexer.IDENT, "c", 0, 11, 3},
TestToken{lexer.DOT, ".", 0, 11, 4},
TestToken{lexer.IDENT, "String", 0, 11, 5},
TestToken{lexer.LPAREN, "(", 0, 11, 11},
TestToken{lexer.IDENT, "arr", 0, 11, 12},
TestToken{lexer.RPAREN, ")", 0, 11, 15},
TestToken{lexer.SEMICOLON, ";", 0, 11, 16},
TestToken{lexer.IDENT, "c", 0, 11, 18},
TestToken{lexer.DOT, ".", 0, 11, 19},
TestToken{lexer.IDENT, "Ln", 0, 11, 20},
TestToken{lexer.SEMICOLON, ";", 0, 11, 22},
TestToken{lexer.IDENT, "NEW", 0, 13, 3},
TestToken{lexer.LPAREN, "(", 0, 13, 6},
TestToken{lexer.IDENT, "fixarr", 0, 13, 7},
TestToken{lexer.RPAREN, ")", 0, 13, 13},
TestToken{lexer.SEMICOLON, ";", 0, 13, 14},
TestToken{lexer.IDENT, "fixarr", 0, 14, 3},
TestToken{lexer.CARAT, "^", 0, 14, 9},
TestToken{lexer.COLONEQUAL, ":=", 0, 14, 11},
TestToken{lexer.STRING, "\"Hello, World!\"", 0, 14, 14},
TestToken{lexer.SEMICOLON, ";", 0, 14, 29},
TestToken{lexer.IDENT, "c", 0, 15, 3},
TestToken{lexer.DOT, ".", 0, 15, 4},
TestToken{lexer.IDENT, "String", 0, 15, 5},
TestToken{lexer.LPAREN, "(", 0, 15, 11},
TestToken{lexer.IDENT, "fixarr", 0, 15, 12},
TestToken{lexer.CARAT, "^", 0, 15, 18},
TestToken{lexer.RPAREN, ")", 0, 15, 19},
TestToken{lexer.SEMICOLON, ";", 0, 15, 20},
TestToken{lexer.IDENT, "c", 0, 15, 22},
TestToken{lexer.DOT, ".", 0, 15, 23},
TestToken{lexer.IDENT, "Ln", 0, 15, 24},
TestToken{lexer.SEMICOLON, ";", 0, 15, 26},
TestToken{lexer.IDENT, "NEW", 0, 17, 3},
TestToken{lexer.LPAREN, "(", 0, 17, 6},
TestToken{lexer.IDENT, "dynarr", 0, 17, 7},
TestToken{lexer.COMMA, ",", 0, 17, 13},
TestToken{lexer.INTEGER, "20", 0, 17, 15},
TestToken{lexer.RPAREN, ")", 0, 17, 17},
TestToken{lexer.SEMICOLON, ";", 0, 17, 18},
TestToken{lexer.IDENT, "dynarr", 0, 18, 3},
TestToken{lexer.CARAT, "^", 0, 18, 9},
TestToken{lexer.COLONEQUAL, ":=", 0, 18, 11},
TestToken{lexer.STRING, "\"Hello, World!\"", 0, 18, 14},
TestToken{lexer.SEMICOLON, ";", 0, 18, 29},
TestToken{lexer.IDENT, "c", 0, 19, 3},
TestToken{lexer.DOT, ".", 0, 19, 4},
TestToken{lexer.IDENT, "String", 0, 19, 5},
TestToken{lexer.LPAREN, "(", 0, 19, 11},
TestToken{lexer.IDENT, "dynarr", 0, 19, 12},
TestToken{lexer.CARAT, "^", 0, 19, 18},
TestToken{lexer.RPAREN, ")", 0, 19, 19},
TestToken{lexer.SEMICOLON, ";", 0, 19, 20},
TestToken{lexer.IDENT, "c", 0, 19, 22},
TestToken{lexer.DOT, ".", 0, 19, 23},
TestToken{lexer.IDENT, "Ln", 0, 19, 24},
TestToken{lexer.SEMICOLON, ";", 0, 19, 26},
TestToken{lexer.END, "END", 0, 20, 1},
TestToken{lexer.IDENT, "Sample2", 0, 20, 5},
TestToken{lexer.DOT, ".", 0, 20, 12},
TestToken{lexer.EOF, "<EOF>", 0, 20, 13},
}
go mylexer.Lexer(s, c)
lexer.CheckTest(s, c, v, t)
}
