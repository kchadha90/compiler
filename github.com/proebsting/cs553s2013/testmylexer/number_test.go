// 4e93b7421a260953d822d103467435c4
package testmylexer

import (
"cs553s2013/mylexer"
"github.com/proebsting/cs553s2013/lexer"
"testing"
)

func Test_4e93b7421a260953d822d103467435c4_20130120105815(t *testing.T) {
c := make(chan lexer.Token, 100)
s := `123H
123ABH
123X
123ABX
123.E+456
123.D+456
123.E-456
123.D-456
123.E456
123.D456
123.89E+456
123.89D+456
123.89E-456
123.89D-456
123.89E456
123.89D456
123.9E+456
123.9D+456
123.9E-456
123.9D-456
123.9E456
123.9D456
123abH
123abh

3H
3ABH
3X
3ABX
3.E+5
3.D+5
3.E-5
3.D-5
3.E5
3.D5
3.89E+5
3.89D+5
3.89E-5
3.89D-5
3.89E5
3.89D5
3.9E+5
3.9D+5
3.9E-5
3.9D-5
3.9E5
3.9D5
3abH
3abh

123H
123ABH
123X
123ABX
123.E+5
123.D+5
123.E-5
123.D-5
123.E5
123.D5
123.89E+5
123.89D+5
123.89E-5
123.89D-5
123.89E5
123.89D5
123.9E+5
123.9D+5
123.9E-5
123.9D-5
123.9E5
123.9D5
123abH
123abh

3H
3ABH
3X
3ABX
3.E+5
3.D+5
3.E-5
3.D-5
3.E5
3.D5
3.89E+5
3.89D+5
3.89E-5
3.89D-5
3.89E5
3.89D5
3.9E+5
3.9D+5
3.9E-5
3.9D-5
3.9E5
3.9D5
3abH
3abh
`
v := []lexer.Token{
TestToken{lexer.INTEGER, "123H", 0, 1, 1},
TestToken{lexer.INTEGER, "123ABH", 0, 2, 1},
TestToken{lexer.STRING, "123X", 0, 3, 1},
TestToken{lexer.STRING, "123ABX", 0, 4, 1},
TestToken{lexer.REAL, "123.E+456", 0, 5, 1},
TestToken{lexer.REAL, "123.D+456", 0, 6, 1},
TestToken{lexer.REAL, "123.E-456", 0, 7, 1},
TestToken{lexer.REAL, "123.D-456", 0, 8, 1},
TestToken{lexer.REAL, "123.E456", 0, 9, 1},
TestToken{lexer.REAL, "123.D456", 0, 10, 1},
TestToken{lexer.REAL, "123.89E+456", 0, 11, 1},
TestToken{lexer.REAL, "123.89D+456", 0, 12, 1},
TestToken{lexer.REAL, "123.89E-456", 0, 13, 1},
TestToken{lexer.REAL, "123.89D-456", 0, 14, 1},
TestToken{lexer.REAL, "123.89E456", 0, 15, 1},
TestToken{lexer.REAL, "123.89D456", 0, 16, 1},
TestToken{lexer.REAL, "123.9E+456", 0, 17, 1},
TestToken{lexer.REAL, "123.9D+456", 0, 18, 1},
TestToken{lexer.REAL, "123.9E-456", 0, 19, 1},
TestToken{lexer.REAL, "123.9D-456", 0, 20, 1},
TestToken{lexer.REAL, "123.9E456", 0, 21, 1},
TestToken{lexer.REAL, "123.9D456", 0, 22, 1},
TestToken{lexer.INTEGER, "123", 0, 23, 1},
TestToken{lexer.IDENT, "abH", 0, 23, 4},
TestToken{lexer.INTEGER, "123", 0, 24, 1},
TestToken{lexer.IDENT, "abh", 0, 24, 4},
TestToken{lexer.INTEGER, "3H", 0, 26, 1},
TestToken{lexer.INTEGER, "3ABH", 0, 27, 1},
TestToken{lexer.STRING, "3X", 0, 28, 1},
TestToken{lexer.STRING, "3ABX", 0, 29, 1},
TestToken{lexer.REAL, "3.E+5", 0, 30, 1},
TestToken{lexer.REAL, "3.D+5", 0, 31, 1},
TestToken{lexer.REAL, "3.E-5", 0, 32, 1},
TestToken{lexer.REAL, "3.D-5", 0, 33, 1},
TestToken{lexer.REAL, "3.E5", 0, 34, 1},
TestToken{lexer.REAL, "3.D5", 0, 35, 1},
TestToken{lexer.REAL, "3.89E+5", 0, 36, 1},
TestToken{lexer.REAL, "3.89D+5", 0, 37, 1},
TestToken{lexer.REAL, "3.89E-5", 0, 38, 1},
TestToken{lexer.REAL, "3.89D-5", 0, 39, 1},
TestToken{lexer.REAL, "3.89E5", 0, 40, 1},
TestToken{lexer.REAL, "3.89D5", 0, 41, 1},
TestToken{lexer.REAL, "3.9E+5", 0, 42, 1},
TestToken{lexer.REAL, "3.9D+5", 0, 43, 1},
TestToken{lexer.REAL, "3.9E-5", 0, 44, 1},
TestToken{lexer.REAL, "3.9D-5", 0, 45, 1},
TestToken{lexer.REAL, "3.9E5", 0, 46, 1},
TestToken{lexer.REAL, "3.9D5", 0, 47, 1},
TestToken{lexer.INTEGER, "3", 0, 48, 1},
TestToken{lexer.IDENT, "abH", 0, 48, 2},
TestToken{lexer.INTEGER, "3", 0, 49, 1},
TestToken{lexer.IDENT, "abh", 0, 49, 2},
TestToken{lexer.INTEGER, "123H", 0, 51, 1},
TestToken{lexer.INTEGER, "123ABH", 0, 52, 1},
TestToken{lexer.STRING, "123X", 0, 53, 1},
TestToken{lexer.STRING, "123ABX", 0, 54, 1},
TestToken{lexer.REAL, "123.E+5", 0, 55, 1},
TestToken{lexer.REAL, "123.D+5", 0, 56, 1},
TestToken{lexer.REAL, "123.E-5", 0, 57, 1},
TestToken{lexer.REAL, "123.D-5", 0, 58, 1},
TestToken{lexer.REAL, "123.E5", 0, 59, 1},
TestToken{lexer.REAL, "123.D5", 0, 60, 1},
TestToken{lexer.REAL, "123.89E+5", 0, 61, 1},
TestToken{lexer.REAL, "123.89D+5", 0, 62, 1},
TestToken{lexer.REAL, "123.89E-5", 0, 63, 1},
TestToken{lexer.REAL, "123.89D-5", 0, 64, 1},
TestToken{lexer.REAL, "123.89E5", 0, 65, 1},
TestToken{lexer.REAL, "123.89D5", 0, 66, 1},
TestToken{lexer.REAL, "123.9E+5", 0, 67, 1},
TestToken{lexer.REAL, "123.9D+5", 0, 68, 1},
TestToken{lexer.REAL, "123.9E-5", 0, 69, 1},
TestToken{lexer.REAL, "123.9D-5", 0, 70, 1},
TestToken{lexer.REAL, "123.9E5", 0, 71, 1},
TestToken{lexer.REAL, "123.9D5", 0, 72, 1},
TestToken{lexer.INTEGER, "123", 0, 73, 1},
TestToken{lexer.IDENT, "abH", 0, 73, 4},
TestToken{lexer.INTEGER, "123", 0, 74, 1},
TestToken{lexer.IDENT, "abh", 0, 74, 4},
TestToken{lexer.INTEGER, "3H", 0, 76, 1},
TestToken{lexer.INTEGER, "3ABH", 0, 77, 1},
TestToken{lexer.STRING, "3X", 0, 78, 1},
TestToken{lexer.STRING, "3ABX", 0, 79, 1},
TestToken{lexer.REAL, "3.E+5", 0, 80, 1},
TestToken{lexer.REAL, "3.D+5", 0, 81, 1},
TestToken{lexer.REAL, "3.E-5", 0, 82, 1},
TestToken{lexer.REAL, "3.D-5", 0, 83, 1},
TestToken{lexer.REAL, "3.E5", 0, 84, 1},
TestToken{lexer.REAL, "3.D5", 0, 85, 1},
TestToken{lexer.REAL, "3.89E+5", 0, 86, 1},
TestToken{lexer.REAL, "3.89D+5", 0, 87, 1},
TestToken{lexer.REAL, "3.89E-5", 0, 88, 1},
TestToken{lexer.REAL, "3.89D-5", 0, 89, 1},
TestToken{lexer.REAL, "3.89E5", 0, 90, 1},
TestToken{lexer.REAL, "3.89D5", 0, 91, 1},
TestToken{lexer.REAL, "3.9E+5", 0, 92, 1},
TestToken{lexer.REAL, "3.9D+5", 0, 93, 1},
TestToken{lexer.REAL, "3.9E-5", 0, 94, 1},
TestToken{lexer.REAL, "3.9D-5", 0, 95, 1},
TestToken{lexer.REAL, "3.9E5", 0, 96, 1},
TestToken{lexer.REAL, "3.9D5", 0, 97, 1},
TestToken{lexer.INTEGER, "3", 0, 98, 1},
TestToken{lexer.IDENT, "abH", 0, 98, 2},
TestToken{lexer.INTEGER, "3", 0, 99, 1},
TestToken{lexer.IDENT, "abh", 0, 99, 2},
TestToken{lexer.EOF, "<EOF>", 0, 100, 1},
}
go mylexer.Lexer(s, c)
lexer.CheckTest(s, c, v, t)
}
