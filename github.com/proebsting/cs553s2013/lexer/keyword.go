package lexer

func Keyword(s string) (token int, match bool) {
	switch s {
	case "ARRAY":
		return ARRAY, true
	case "BEGIN":
		return BEGIN, true
	case "BY":
		return BY, true
	case "CASE":
		return CASE, true
	case "CONST":
		return CONST, true
	case "DIV":
		return DIV, true
	case "DO":
		return DO, true
	case "ELSE":
		return ELSE, true
	case "ELSIF":
		return ELSIF, true
	case "END":
		return END, true
	case "FALSE":
		return FALSE, true
	case "FOR":
		return FOR, true
	case "IF":
		return IF, true
	case "IMPORT":
		return IMPORT, true
	case "IN":
		return IN, true
	case "IS":
		return IS, true
	case "MOD":
		return MOD, true
	case "MODULE":
		return MODULE, true
	case "NIL":
		return NIL, true
	case "OF":
		return OF, true
	case "OR":
		return OR, true
	case "POINTER":
		return POINTER, true
	case "PROCEDURE":
		return PROCEDURE, true
	case "RECORD":
		return RECORD, true
	case "REPEAT":
		return REPEAT, true
	case "RETURN":
		return RETURN, true
	case "THEN":
		return THEN, true
	case "TO":
		return TO, true
	case "TRUE":
		return TRUE, true
	case "TYPE":
		return TYPE, true
	case "UNTIL":
		return UNTIL, true
	case "VAR":
		return VAR, true
	case "WHILE":
		return WHILE, true
	}
	return -1, false
}
