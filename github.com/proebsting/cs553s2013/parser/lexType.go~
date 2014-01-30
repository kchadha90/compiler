package parser

import (
"github.com/proebsting/cs553s2013/lexer"
)

type LexType struct {
stream chan lexer.Token
current lexer.Token
}

func (o *LexType) Peek() (val int) {
return o.current.Enum()
}

func (o *LexType) Advance() {
o.current = <-o.stream
}

func (o *LexType) ErrorMessage(s string) (val string) {
return s
}
