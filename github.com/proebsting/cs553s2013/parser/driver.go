package parser

import (
"cs553s2013/mylexer"
//"fmt"
"github.com/proebsting/cs553s2013/lexer"
)

func Driver(input string)(bool,string) {
ch := make(chan lexer.Token, 100)
go mylexer.Lexer(input, ch)
v := <-ch
lex := LexType{ch, v}
ok, msg := Module(&lex)
return ok,msg
}
