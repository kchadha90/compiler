package lexer

func Operator(s string) (token int, length int, match bool) {
	if len(s) == 0 {
		return -1, -1, false
	}
	switch s[0] {
	case '#':
		return HASH, 1, true
	case '&':
		return AMP, 1, true
	case '(':
		return LPAREN, 1, true
	case ')':
		return RPAREN, 1, true
	case '*':
		return STAR, 1, true
	case '+':
		return PLUS, 1, true
	case ',':
		return COMMA, 1, true
	case '-':
		return MINUS, 1, true
	case '.':
		if len(s) > 1 && s[1] == '.' {
			return DOTDOT, 2, true
		} else {
			return DOT, 1, true
		}
	case '/':
		return SLASH, 1, true
	case ':':
		if len(s) > 1 && s[1] == '=' {
			return COLONEQUAL, 2, true
		} else {
			return COLON, 1, true
		}
	case ';':
		return SEMICOLON, 1, true
	case '<':
		if len(s) > 1 && s[1] == '=' {
			return LESSEQUAL, 2, true
		} else {
			return LESS, 1, true
		}
	case '=':
		return EQUAL, 1, true
	case '>':
		if len(s) > 1 && s[1] == '=' {
			return GREATEREQUAL, 2, true
		} else {
			return GREATER, 1, true
		}
	case '[':
		return LSQUARE, 1, true
	case ']':
		return RSQUARE, 1, true
	case '^':
		return CARAT, 1, true
	case '{':
		return LCURLY, 1, true
	case '|':
		return BAR, 1, true
	case '}':
		return RCURLY, 1, true
	case '~':
		return TILDE, 1, true
	}
	return -1, -1, false
}
