package parser

import (
	//"fmt"
	"github.com/proebsting/cs553s2013/lexer"
	"strconv"
)

type MODULE struct { //n1
	mid  *IDENT
	dseq *DECSEQ
	sseq *STATSEQ
	eid  *IDENT
}

type DECSEQ struct { //n2
	tdec []*TDEC
	vdec []*VDEC
	pdec *PDEC
}

type PDEC struct { //n26
	phead *PHEAD
	pbody *PBODY
	id    *IDENT
}

type PHEAD struct { //n27
	id    *IDENT
	fpsec []*FPSEC
	id2   *IDENT
}

type PBODY struct { //n3
	decseq  *DECSEQ
	statseq *STATSEQ
	exp     *EXP
}

type TDEC struct { //n20
	id *IDENT
	at *AT
}

type AT struct { //21
	integer string
	ty      *TYPE
}

type TYPE struct { //n22
	id *IDENT
}

type STATSEQ struct { //n4
	statmnt []*STATMNT
}

type STATMNT struct { //n8
	aoc    *AOC
	ifstat *IFSTAT
	wstat  *WSTAT
	fstat  *FSTAT
}

type AOC struct { //n13
	desg *DESG
	exp  *EXP
	ap   *AP
}

type IFSTAT struct { //n5
	ifexp     *EXP
	statseq   *STATSEQ
	elseifexp *EXP
	eistatseq *STATSEQ
	estatseq  *STATSEQ
}

type WSTAT struct { //n6
	exp     *EXP
	statseq *STATSEQ
}

type FSTAT struct { //n7
	id      *IDENT
	exp1    *EXP
	exp2    *EXP
	value   string
	statseq *STATSEQ
}

type DESG struct { //n17
	id  *IDENT
	sel *SEL
}

type EXP struct { //n9
	se  *SE
	re  *RE
	se2 *SE
}

type VDEC struct { //n23
	Il *IDENTLIST
	ty *TYPE
}

type AP struct { //n14
	expl *EXPLIST
}

type EXPLIST struct { //n12
	e []*EXP
}

type SEL struct { //n11
	exp *EXP
}

type SE struct { //n16
	op    string
	term  *TERM
	addop *ADDOP
	term2 *TERM
}

type RE struct { //n25
	reop string
}

type TERM struct { //n15
	fact  *FACT
	mulop *MULOP
	fact2 *FACT
}

type ADDOP struct { //n18
	op string
}

type MULOP struct { //n19
	op string
}

type FACT struct { //n10
	integer *IDENT
	str     *IDENT
	bol1    *IDENT
	bol2    *IDENT
	dsg     *DESG
	ap      *AP
	exp     *EXP
	til     *IDENT
	fac     *FACT
}

type FPSEC struct { //n28
	id    []*IDENT
	ftype *FTYPE
}

type FTYPE struct { //n29
	id *IDENT
}

type IDENTLIST struct { //n24
	idl []*IDENT
	//id2 *IDENT
}

type IDENT struct { //n0
	id     string
	ty     int
	lineno int
}

func TypeDeclaration(lex *LexType) (n *TDEC, ok bool, error string) {
	n20 := new(TDEC)
	n0 := new(IDENT)
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n20.id = n0
		lex.Advance()
	} else {
		return n20, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.EQUAL == lex.Peek() {
		lex.Advance()
	} else {
		return n20, false, lex.ErrorMessage(" expected = at line number " + strconv.Itoa(lex.current.Line()))
	}
	n21, ok, msg := ArrayType(lex)
	n20.at = n21
	if !ok {
		return n20, ok, msg
	}
	return n20, true, "ok"
}

func Type(lex *LexType) (n *TYPE, ok bool, error string) {
	n22 := new(TYPE)
	n0 := new(IDENT)
	switch lex.Peek() {
	case lexer.IDENT:
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n22.id = n0
		lex.Advance() // IDENT
	/*case lexer.ARRAY:
	n21, ok, msg := ArrayType(lex);
	if !ok {
		return ok, msg
	}*/
	default:
		return n22, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n22, true, "ok"
}

func ArrayType(lex *LexType) (n *AT, ok bool, error string) {
	n21 := new(AT)
	if lexer.ARRAY == lex.Peek() {
		lex.Advance()
	} else {
		return n21, false, lex.ErrorMessage(" expected ARRAY at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.INTEGER == lex.Peek() {
		n21.integer = lex.current.Value()
		lex.Advance()
	} else {
		return n21, false, lex.ErrorMessage(" expected integer at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.OF == lex.Peek() {
		lex.Advance()
	} else {
		return n21, false, lex.ErrorMessage(" expected OF at line number " + strconv.Itoa(lex.current.Line()))
	}
	n22, ok, msg := Type(lex)
	n21.ty = n22
	if !ok {
		return n21, ok, msg
	}
	return n21, true, "ok"
}

func IdentList(lex *LexType) (n *IDENTLIST, ok bool, error string) {
	n24 := new(IDENTLIST)
	n0 := new(IDENT)
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.ty = lexer.IDENT
		n0.lineno = lex.current.Line()
		n24.idl = append(n24.idl, n0)
		//fmt.Println(n24.idl)
		lex.Advance()
	} else {
		return n24, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
L0:
	for {
		switch lex.Peek() {
		case lexer.COMMA:
			lex.Advance() // COMMA
			if lexer.IDENT == lex.Peek() {
				d0 := new(IDENT)
				d0.id = lex.current.Value()
				d0.ty = lexer.IDENT
				d0.lineno = lex.current.Line()
				n24.idl = append(n24.idl, d0)
				//fmt.Println(n24.idl)
				lex.Advance()
			} else {
				return n24, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
			}
		default:
			break L0
		}
	}
	return n24, true, "ok"
}

func VariableDeclaration(lex *LexType) (n *VDEC, ok bool, error string) {
	n23 := new(VDEC)
	n24, ok, msg := IdentList(lex)
	n23.Il = n24
	if !ok {
		return n23, ok, msg
	}
	if lexer.COLON == lex.Peek() {
		lex.Advance()
	} else {
		return n23, false, lex.ErrorMessage(" expected : at line number " + strconv.Itoa(lex.current.Line()))
	}
	n22, ok, msg := Type(lex)
	n23.ty = n22
	if !ok {
		return n23, ok, msg
	}
	return n23, true, "ok"
}

func expression(lex *LexType) (n *EXP, ok bool, error string) {
	n9 := new(EXP)
	n16, ok, msg := SimpleExpression(lex)
	n9.se = n16
	if !ok {
		return n9, ok, msg
	}
	switch lex.Peek() {
	case lexer.HASH, lexer.GREATEREQUAL, lexer.EQUAL, lexer.GREATER, lexer.LESS, lexer.LESSEQUAL:
		n25, ok, msg := relation(lex)
		n9.re = n25
		if !ok {
			return n9, ok, msg
		}
		n16, ok, msg := SimpleExpression(lex)
		n9.se2 = n16
		if !ok {
			return n9, ok, msg
		}
	}
	return n9, true, "ok"
}

func relation(lex *LexType) (n *RE, ok bool, error string) {
	n25 := new(RE)
	switch lex.Peek() {
	case lexer.EQUAL:
		n25.reop = lex.current.Value()
		lex.Advance() // EQUAL
	case lexer.HASH:
		n25.reop = lex.current.Value()
		lex.Advance() // HASH
	case lexer.LESS:
		n25.reop = lex.current.Value()
		lex.Advance() // LESS
	case lexer.LESSEQUAL:
		n25.reop = lex.current.Value()
		lex.Advance() // LESSEQUAL
	case lexer.GREATER:
		n25.reop = lex.current.Value()
		lex.Advance() // GREATER
	case lexer.GREATEREQUAL:
		n25.reop = lex.current.Value()
		lex.Advance() // GREATEREQUAL
	default:
		return n25, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n25, true, "ok"
}

func SimpleExpression(lex *LexType) (n *SE, ok bool, error string) {
	n16 := new(SE)
	switch lex.Peek() {
	case lexer.MINUS, lexer.PLUS:
		switch lex.Peek() {
		case lexer.PLUS:
			n16.op = "+"
			lex.Advance() // PLUS
		case lexer.MINUS:
			n16.op = "-"
			lex.Advance() // MINUS
		default:
			return n16, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
		}
	}
	n15, ok, msg := term(lex)
	n16.term = n15
	if !ok {
		return n16, ok, msg
	}
L1:
	for {
		switch lex.Peek() {
		case lexer.MINUS, lexer.OR, lexer.PLUS:
			n19, ok, msg := AddOperator(lex)
			n16.addop = n19
			if !ok {
				return n16, ok, msg
			}
			n15, ok, msg := term(lex)
			n16.term2 = n15
			if !ok {
				return n16, ok, msg
			}
		default:
			break L1
		}
	}
	return n16, true, "ok"
}

func AddOperator(lex *LexType) (n *ADDOP, ok bool, error string) {
	n19 := new(ADDOP)
	switch lex.Peek() {
	case lexer.PLUS:
		n19.op = lex.current.Value()
		lex.Advance() // PLUS
	case lexer.MINUS:
		n19.op = lex.current.Value()
		lex.Advance() // MINUS
	case lexer.OR:
		n19.op = lex.current.Value()
		lex.Advance() // OR
	default:
		return n19, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n19, true, "ok"
}

func term(lex *LexType) (n *TERM, ok bool, error string) {
	n15 := new(TERM)
	n10, ok, msg := factor(lex)
	n15.fact = n10
	if !ok {
		return n15, ok, msg
	}
L2:
	for {
		switch lex.Peek() {
		case lexer.STAR, lexer.DIV, lexer.MOD, lexer.SLASH, lexer.AMP:
			n18, ok, msg := MulOperator(lex)
			n15.mulop = n18
			//fmt.Println(n15.mulop)
			if !ok {
				return n15, ok, msg
			}
			n10, ok, msg := factor(lex)
			n15.fact2 = n10
			//fmt.Println(n15.fact2)
			if !ok {
				return n15, ok, msg
			}
		default:
			break L2
		}
	}
	//fmt.Println(n15.mulop)
	return n15, true, "ok"
}

func MulOperator(lex *LexType) (n *MULOP, ok bool, error string) {
	n18 := new(MULOP)
	switch lex.Peek() {
	case lexer.STAR:
		n18.op = lex.current.Value()
		//fmt.Println(n18)
		lex.Advance() // STAR
	case lexer.SLASH:
		n18.op = lex.current.Value()
		lex.Advance() // SLASH
	case lexer.DIV:
		n18.op = lex.current.Value()
		lex.Advance() // DIV
	case lexer.MOD:
		n18.op = lex.current.Value()
		lex.Advance() // MOD
	case lexer.AMP:
		n18.op = lex.current.Value()
		lex.Advance() // AMP
	default:
		return n18, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n18, true, "ok"
}

func factor(lex *LexType) (n *FACT, ok bool, error string) {
	n10 := new(FACT)
	n0 := new(IDENT)
	switch lex.Peek() {
	case lexer.INTEGER:
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n10.integer = n0
		lex.Advance() // INTEGER
	case lexer.STRING:
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n10.str = n0
		lex.Advance() // STRING
	/*case lexer.NIL:
	n10.n = "NIL"
	lex.Advance() // NIL*/
	case lexer.TRUE:
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n10.bol1 = n0
		lex.Advance() // TRUE
	case lexer.FALSE:
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n10.bol2 = n0
		lex.Advance() // FALSE
	case lexer.IDENT:
		n17, ok, msg := designator(lex)
		n10.dsg = n17
		if !ok {
			return n10, ok, msg
		}
		switch lex.Peek() {
		case lexer.LPAREN:
			n14, ok, msg := ActualParameters(lex)
			n10.ap = n14
			if !ok {
				return n10, ok, msg
			}
		}
	case lexer.LPAREN:
		lex.Advance() // LPAREN
		n9, ok, msg := expression(lex)
		n10.exp = n9
		if !ok {
			return n10, ok, msg
		}
		if lexer.RPAREN == lex.Peek() {
			lex.Advance()
		} else {
			return n10, false, lex.ErrorMessage(" expected ) at line number " + strconv.Itoa(lex.current.Line()))
		}
	case lexer.TILDE:
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n10.til = n0
		lex.Advance() // TILDE
		n10, ok, msg := factor(lex)
		n10.fac = n10
		if !ok {
			return n10, ok, msg
		}
	default:
		return n10, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n10, true, "ok"
}

func designator(lex *LexType) (n *DESG, ok bool, error string) {
	n17 := new(DESG)
	n0 := new(IDENT)
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n17.id = n0
		lex.Advance()
	} else {
		return n17, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
L3:
	for {
		switch lex.Peek() {
		case lexer.LSQUARE:
			n11, ok, msg := selector(lex)
			n17.sel = n11
			if !ok {
				return n17, ok, msg
			}
		default:
			break L3
		}
	}
	return n17, true, "ok"
}

func selector(lex *LexType) (n *SEL, ok bool, error string) {
	n11 := new(SEL)
	if lexer.LSQUARE == lex.Peek() {
		lex.Advance()
	} else {
		return n11, false, lex.ErrorMessage(" expected [ at line number " + strconv.Itoa(lex.current.Line()))
	}
	n9, ok, msg := expression(lex)
	n11.exp = n9
	if !ok {
		return n11, ok, msg
	}
	if lexer.RSQUARE == lex.Peek() {
		lex.Advance()
	} else {
		return n11, false, lex.ErrorMessage(" expected ] at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n11, true, "ok"
}

func ExpList(lex *LexType) (n *EXPLIST, ok bool, error string) {
	n12 := new(EXPLIST)
	n9, ok, msg := expression(lex)
	n12.e = append(n12.e, n9)
	if !ok {
		return n12, ok, msg
	}
L4:
	for {
		switch lex.Peek() {
		case lexer.COMMA:
			lex.Advance() // COMMA
			n9, ok, msg := expression(lex)
			n12.e = append(n12.e, n9)
			if !ok {
				return n12, ok, msg
			}
		default:
			break L4
		}
	}
	return n12, true, "ok"
}

func ActualParameters(lex *LexType) (n *AP, ok bool, error string) {
	n14 := new(AP)
	if lexer.LPAREN == lex.Peek() {
		lex.Advance()
	} else {
		return n14, false, lex.ErrorMessage(" expected ( at line number " + strconv.Itoa(lex.current.Line()))
	}
	switch lex.Peek() {
	case lexer.MINUS, lexer.FALSE, lexer.INTEGER, lexer.NIL, lexer.TILDE, lexer.PLUS, lexer.TRUE, lexer.LPAREN, lexer.IDENT, lexer.STRING:
		n12, ok, msg := ExpList(lex)
		n14.expl = n12
		if !ok {
			return n14, ok, msg
		}
	}
	if lexer.RPAREN == lex.Peek() {
		lex.Advance()
	} else {
		return n14, false, lex.ErrorMessage(" expected ) at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n14, true, "ok"
}

func statement(lex *LexType) (n *STATMNT, ok bool, error string) {

	n8 := new(STATMNT)
	switch lex.Peek() {
	case lexer.IF, lexer.FOR, lexer.IDENT, lexer.WHILE:
		switch lex.Peek() {
		case lexer.IDENT:
			n13, ok, msg := AssignOrCall(lex)
			n8.aoc = n13
			if !ok {
				return n8, ok, msg
			}
		case lexer.IF:
			n5, ok, msg := IfStatement(lex)
			n8.ifstat = n5
			if !ok {
				return n8, ok, msg
			}
		case lexer.WHILE:
			n6, ok, msg := WhileStatement(lex)
			n8.wstat = n6
			if !ok {
				return n8, ok, msg
			}
		case lexer.FOR:
			n7, ok, msg := ForStatement(lex)
			n8.fstat = n7
			if !ok {
				return n8, ok, msg
			}
		default:
			return n8, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
		}
	}
	return n8, true, "ok"
}

func AssignOrCall(lex *LexType) (n *AOC, ok bool, error string) {
	n13 := new(AOC)
	n17, ok, msg := designator(lex)
	n13.desg = n17
	if !ok {
		return n13, ok, msg
	}
	switch lex.Peek() {
	case lexer.COLONEQUAL:
		lex.Advance() // COLONEQUAL
		n9, ok, msg := expression(lex)
		n13.exp = n9
		if !ok {
			return n13, ok, msg
		}
	case lexer.LPAREN:
		n14, ok, msg := ActualParameters(lex)
		n13.ap = n14
		if !ok {
			return n13, ok, msg
		}
	default:
		return n13, false, lex.ErrorMessage(" unexpected token at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n13, true, "ok"
}

func StatementSequence(lex *LexType) (n *STATSEQ, ok bool, error string) {
	n4 := new(STATSEQ)
	n8, ok, msg := statement(lex)
	n4.statmnt = append(n4.statmnt, n8)
	if !ok {
		return n4, ok, msg
	}
L5:
	for {
		switch lex.Peek() {
		case lexer.SEMICOLON:
			lex.Advance() // SEMICOLON
			n8, ok, msg := statement(lex)
			n4.statmnt = append(n4.statmnt, n8)
			if !ok {
				return n4, ok, msg
			}
		default:
			break L5
		}
	}
	return n4, true, "ok"
}

func IfStatement(lex *LexType) (n *IFSTAT, ok bool, error string) {
	n5 := new(IFSTAT)
	if lexer.IF == lex.Peek() {
		lex.Advance()
	} else {
		return n5, false, lex.ErrorMessage(" expected IF at line number " + strconv.Itoa(lex.current.Line()))
	}
	n9, ok, msg := expression(lex)
	n5.ifexp = n9
	if !ok {
		return n5, ok, msg
	}
	if lexer.THEN == lex.Peek() {
		lex.Advance()
	} else {
		return n5, false, lex.ErrorMessage(" expected THEN at line number " + strconv.Itoa(lex.current.Line()))
	}
	n4, ok, msg := StatementSequence(lex)
	n5.statseq = n4
	if !ok {
		return n5, ok, msg
	}
L6:
	for {
		switch lex.Peek() {
		case lexer.ELSIF:
			lex.Advance() // ELSIF
			n9, ok, msg := expression(lex)
			n5.elseifexp = n9
			if !ok {
				return n5, ok, msg
			}
			if lexer.THEN == lex.Peek() {
				lex.Advance()
			} else {
				return n5, false, lex.ErrorMessage(" expected THEN at line number " + strconv.Itoa(lex.current.Line()))
			}
			n4, ok, msg := StatementSequence(lex)
			n5.eistatseq = n4
			if !ok {
				return n5, ok, msg
			}
		default:
			break L6
		}
	}
	switch lex.Peek() {
	case lexer.ELSE:
		lex.Advance() // ELSE
		n4, ok, msg := StatementSequence(lex)
		n5.estatseq = n4
		if !ok {
			return n5, ok, msg
		}
	}
	if lexer.END == lex.Peek() {
		lex.Advance()
	} else {
		return n5, false, lex.ErrorMessage(" expected END at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n5, true, "ok"
}

func WhileStatement(lex *LexType) (n *WSTAT, ok bool, error string) {

	n6 := new(WSTAT)

	if lexer.WHILE == lex.Peek() {
		lex.Advance()
	} else {
		return n6, false, lex.ErrorMessage(" expected WHILE at line number " + strconv.Itoa(lex.current.Line()))
	}
	n9, ok, msg := expression(lex)
	n6.exp = n9
	if !ok {
		return n6, ok, msg
	}
	if lexer.DO == lex.Peek() {
		lex.Advance()
	} else {
		return n6, false, lex.ErrorMessage(" expected DO at line number " + strconv.Itoa(lex.current.Line()))
	}
	n4, ok, msg := StatementSequence(lex)
	n6.statseq = n4
	if !ok {
		return n6, ok, msg
	}
	/*L7:
	for {
		switch lex.Peek() {
		case lexer.ELSIF:
			lex.Advance() // ELSIF
			n9,ok, msg := expression(lex); !ok {
				return ok, msg
			}
			if lexer.DO == lex.Peek() {
				lex.Advance()
			} else {
				return false, lex.ErrorMessage(" expected DO")
			}
			n4,ok, msg := StatementSequence(lex);
			n5.statseq = n4
			if !ok {
				return ok, msg
			}
		default:
			break L7
		}
	}*/
	if lexer.END == lex.Peek() {
		lex.Advance()
	} else {
		return n6, false, lex.ErrorMessage(" expected END at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n6, true, "ok"
}

func ForStatement(lex *LexType) (n *FSTAT, ok bool, error string) {
	n7 := new(FSTAT)
	n0 := new(IDENT)
	if lexer.FOR == lex.Peek() {
		lex.Advance()
	} else {
		return n7, false, lex.ErrorMessage(" expected FOR at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n7.id = n0
		lex.Advance()
	} else {
		return n7, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.COLONEQUAL == lex.Peek() {
		lex.Advance()
	} else {
		return n7, false, lex.ErrorMessage(" expected := at line number " + strconv.Itoa(lex.current.Line()))
	}
	n9, ok, msg := expression(lex)
	n7.exp1 = n9
	if !ok {
		return n7, ok, msg
	}
	if lexer.TO == lex.Peek() {
		lex.Advance()
	} else {
		return n7, false, lex.ErrorMessage(" expected TO at line number " + strconv.Itoa(lex.current.Line()))
	}
	n9_2, ok, msg := expression(lex)
	n7.exp2 = n9_2
	if !ok {
		return n7, ok, msg
	}
	switch lex.Peek() {
	case lexer.BY:
		lex.Advance() // BY
		if lexer.INTEGER == lex.Peek() {
			n7.value = lex.current.Value()
			lex.Advance()
		} else {
			return n7, false, lex.ErrorMessage(" expected integer at line number " + strconv.Itoa(lex.current.Line()))
		}
	}
	if lexer.DO == lex.Peek() {
		lex.Advance()
	} else {
		return n7, false, lex.ErrorMessage(" expected DO at line number " + strconv.Itoa(lex.current.Line()))
	}
	n4, ok, msg := StatementSequence(lex)
	n7.statseq = n4
	if !ok {
		return n7, ok, msg
	}
	if lexer.END == lex.Peek() {
		lex.Advance()
	} else {
		return n7, false, lex.ErrorMessage(" expected END at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n7, true, "ok"
}

func ProcedureDeclaration(lex *LexType) (n *PDEC, ok bool, error string) {
	n26 := new(PDEC)
	n0 := new(IDENT)
	n27, ok, msg := ProcedureHeading(lex)
	n26.phead = n27
	if !ok {
		return n26, ok, msg
	}
	if lexer.SEMICOLON == lex.Peek() {
		lex.Advance()
	} else {
		return n26, false, lex.ErrorMessage(" expected ; at line number " + strconv.Itoa(lex.current.Line()))
	}
	n3, ok, msg := ProcedureBody(lex)
	n26.pbody = n3
	if !ok {
		return n26, ok, msg
	}
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n26.id = n0
		lex.Advance()
	} else {
		return n26, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n26, true, "ok"
}

func ProcedureHeading(lex *LexType) (n *PHEAD, ok bool, error string) {
	n27 := new(PHEAD)
	n0 := new(IDENT)
	if lexer.PROCEDURE == lex.Peek() {
		lex.Advance()
	} else {
		return n27, false, lex.ErrorMessage(" expected PROCEDURE at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n27.id = n0
		lex.Advance()
	} else {
		return n27, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.LPAREN == lex.Peek() {
		lex.Advance()
	} else {
		return n27, false, lex.ErrorMessage(" expected ( at line number " + strconv.Itoa(lex.current.Line()))
	}
	switch lex.Peek() {
	case lexer.VAR, lexer.IDENT:
		n28, ok, msg := FPSection(lex)
		n27.fpsec = append(n27.fpsec, n28)
		if !ok {
			return n27, ok, msg
		}
	L8:
		for {
			switch lex.Peek() {
			case lexer.SEMICOLON:
				lex.Advance() // SEMICOLON
				n28, ok, msg := FPSection(lex)
				n27.fpsec = append(n27.fpsec, n28)
				if !ok {
					return n27, ok, msg
				}
			default:
				break L8
			}
		}
	}
	if lexer.RPAREN == lex.Peek() {
		lex.Advance()
	} else {
		return n27, false, lex.ErrorMessage(" expected ) at line number " + strconv.Itoa(lex.current.Line()))
	}
	switch lex.Peek() {
	case lexer.COLON:
		lex.Advance() // COLON
		if lexer.IDENT == lex.Peek() {
			d0 := new(IDENT)
			d0.id = lex.current.Value()
			d0.lineno = lex.current.Line()
			n27.id2 = d0
			lex.Advance()
		} else {
			return n27, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
		}
	}
	return n27, true, "ok"
}

func ProcedureBody(lex *LexType) (n *PBODY, ok bool, error string) {

	n3 := new(PBODY)
	n2, ok, msg := DeclarationSequence(lex)
	n3.decseq = n2
	if !ok {
		return n3, ok, msg
	}

	n3.decseq = n2
	switch lex.Peek() {
	case lexer.BEGIN:
		lex.Advance() // BEGIN
		n4, ok, msg := StatementSequence(lex)
		n3.statseq = n4
		if !ok {
			return n3, ok, msg
		}
	}
	switch lex.Peek() {
	case lexer.RETURN:
		lex.Advance() // RETURN
		n9, ok, msg := expression(lex)
		n3.exp = n9
		if !ok {
			return n3, ok, msg
		}
	}
	if lexer.END == lex.Peek() {
		lex.Advance()
	} else {
		return n3, false, lex.ErrorMessage(" expected END at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n3, true, "ok"
}

func DeclarationSequence(lex *LexType) (n *DECSEQ, ok bool, error string) {
	n2 := new(DECSEQ)
	switch lex.Peek() {
	case lexer.TYPE:
		lex.Advance() // TYPE
	L9:
		for {
			switch lex.Peek() {
			case lexer.IDENT:
				n20, ok, msg := TypeDeclaration(lex)
				n2.tdec = append(n2.tdec, n20)
				if !ok {
					return n2, ok, msg
				}
				if lexer.SEMICOLON == lex.Peek() {
					lex.Advance()
				} else {
					return n2, false, lex.ErrorMessage(" expected ; at line number " + strconv.Itoa(lex.current.Line()))
				}
			default:
				break L9
			}
		}
	}
	switch lex.Peek() {
	case lexer.VAR:
		lex.Advance() // VAR
	L10:
		for {
			switch lex.Peek() {
			case lexer.IDENT:
				n23, ok, msg := VariableDeclaration(lex)
				n2.vdec = append(n2.vdec, n23)
				if !ok {
					return n2, ok, msg
				}
				if lexer.SEMICOLON == lex.Peek() {
					lex.Advance()
				} else {
					return n2, false, lex.ErrorMessage(" expected ; at line number " + strconv.Itoa(lex.current.Line()))
				}
			default:
				break L10
			}
		}
	}
L11:
	for {
		switch lex.Peek() {
		case lexer.PROCEDURE:
			n26, ok, msg := ProcedureDeclaration(lex)
			n2.pdec = n26
			if !ok {
				return n2, ok, msg
			}
			if lexer.SEMICOLON == lex.Peek() {
				lex.Advance()
			} else {
				return n2, false, lex.ErrorMessage(" expected ; at line number " + strconv.Itoa(lex.current.Line()))
			}
		default:
			break L11
		}
	}
	return n2, true, "ok"
}

func FPSection(lex *LexType) (n *FPSEC, ok bool, error string) {
	n28 := new(FPSEC)
	n0 := new(IDENT)
	switch lex.Peek() {
	case lexer.VAR:
		lex.Advance() // VAR
	}
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n28.id = append(n28.id, n0)
		lex.Advance()
	} else {
		return n28, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
L12:
	for {
		switch lex.Peek() {
		case lexer.COMMA:
			lex.Advance() // COMMA
			if lexer.IDENT == lex.Peek() {
				lex.Advance()
			} else {
				return n28, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
			}
		default:
			break L12
		}
	}
	if lexer.COLON == lex.Peek() {
		lex.Advance()
	} else {
		return n28, false, lex.ErrorMessage(" expected : at line number " + strconv.Itoa(lex.current.Line()))
	}
	n29, ok, msg := FormalType(lex)
	n28.ftype = n29
	if !ok {
		return n28, ok, msg
	}
	return n28, true, "ok"
}

func FormalType(lex *LexType) (n *FTYPE, ok bool, error string) {
	n29 := new(FTYPE)
	n0 := new(IDENT)
	switch lex.Peek() {
	case lexer.ARRAY:
		lex.Advance() // ARRAY
		if lexer.OF == lex.Peek() {
			lex.Advance()
		} else {
			return n29, false, lex.ErrorMessage(" expected OF at line number " + strconv.Itoa(lex.current.Line()))
		}
	}
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n29.id = n0
		lex.Advance()
	} else {
		return n29, false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
	return n29, true, "ok"
}

func Module(lex *LexType) (ok bool, error string) {
	n1 := new(MODULE)
	n0 := new(IDENT)
	//fmt.Println(lex.current.Value())
	if lexer.MODULE == lex.Peek() {
		lex.Advance()
	} else {	
		return false, lex.ErrorMessage(" expected MODULE at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.IDENT == lex.Peek() {
		n0.id = lex.current.Value()
		n0.lineno = lex.current.Line()
		n0.ty = lexer.IDENT
		n1.mid = n0
		lex.Advance()
	} else {
		return false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.SEMICOLON == lex.Peek() {
		lex.Advance()
	} else {
		return false, lex.ErrorMessage(" expected ; at line number " + strconv.Itoa(lex.current.Line()))
	}
	n2, ok, msg := DeclarationSequence(lex)
	n1.dseq = n2
	if !ok {
		return ok, msg
	}

	switch lex.Peek() {
	case lexer.BEGIN:
		lex.Advance() // BEGIN
		n4, ok, msg := StatementSequence(lex)
		n1.sseq = n4
		if !ok {
			return ok, msg
		}
	}
	if lexer.END == lex.Peek() {
		lex.Advance()
	} else {
		return false, lex.ErrorMessage(" expected END at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.IDENT == lex.Peek() {
		d0 := new(IDENT)
		d0.id = lex.current.Value()
		d0.ty = lexer.IDENT
		d0.lineno = lex.current.Line()
		n1.eid = d0
		lex.Advance()
	} else {
		return false, lex.ErrorMessage(" expected ident at line number " + strconv.Itoa(lex.current.Line()))
	}
	if lexer.DOT == lex.Peek() {
		lex.Advance()
	} else {
		return false, lex.ErrorMessage(" expected dot at line number " + strconv.Itoa(lex.current.Line()))
	}

	ok1, error1 := Scheck(n1, lex)
	return ok1, error1
	//return ok1,error1
}

var index = 0
var index2 = 0

func Scheck(AST *MODULE, lex *LexType) (bool, string) {
	m = make(map[int]M)
	n = make(map[int]N)
	ok1, error1 := ME(AST, m, n, lex)
	return ok1, error1
	//return ok1,error1
}

func ME(AST *MODULE, m map[int]M, n map[int]N, lex *LexType) (bool, string) { //MODULE
	index = 0
	index2 = 0
	var error1 string
	var ok1 bool
	if AST.mid.id != AST.eid.id {
		return false, lex.ErrorMessage(" IDENT MISMATCH.")
	}
	if AST.dseq != nil {
		ok1, error1 = DQ(AST.dseq, m, n, 0, lex) //1
		if ok1 == false {
			return ok1, error1
		}
	}
	if AST.sseq != nil {
		ok1, error1 := SQ(AST.sseq, m, n, 0, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	return true, "OK"
}

func DQ(AST *DECSEQ, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //Declaration sequence	//2
	var error1 string
	var ok1 bool

	for i := 0; i < len(AST.tdec); i++ {
		if AST.tdec[i] != nil {
			ok1, error1 = TD(AST.tdec[i], m, n, scope, lex)
			if ok1 == false {
				return ok1, error1
			}
		}
	}
	for i := 0; i < len(AST.vdec); i++ {
		if AST.vdec[i] != nil {
			ok1, error1 = VD(AST.vdec[i], m, n, scope, lex)
			if ok1 == false {
				return ok1, error1
			}
		}
	}
	if AST.pdec != nil {
		ok, error := PD(AST.pdec, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	return true, "OK"
}

func SQ(AST *STATSEQ, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //STATEMENT SEQUENCE

	var error1 string
	var ok1 bool

	for i := 0; i < len(AST.statmnt); i++ {
		if AST.statmnt[i] != nil {
			ok1, error1 = ST(AST.statmnt[i], m, n, scope, lex)
			if ok1 == false {
				return ok1, error1
			}
		}
	}

	return true, "OK"

}

func ST(AST *STATMNT, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //STATMENT
	var error1 string
	var ok1 bool

	if AST.aoc != nil {
		ok1, error1 = AC(AST.aoc, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	if AST.ifstat != nil {
		ok1, error1 = IT(AST.ifstat, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	if AST.wstat != nil {
		ok1, error1 = WT(AST.wstat, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	if AST.fstat != nil {
		ok1, error1 = FT(AST.fstat, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}

	return true, "OK"
}

func AC(AST *AOC, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //ASSIGN or CALL
	var name, error1 string
	var ok1 bool
	var lno int
	var value string

	if AST.desg != nil {
		ok1, error1, name, lno = DG(AST.desg, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	} //LHS of expression
	if AST.exp != nil {
		ok1, error1, value, lno = EXPR(AST.exp, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}

	}
	if name != "" && value != "" {

		ok2, error2 := Checksemantics(name, value, lno, m, n, scope, lex)
		if ok2 == false {
			return ok2, error2
		}
	} // RHS of expression
	if AST.ap != nil {
		ok1, error1 = APA(AST.ap, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	return true, "OK"
}

func IT(AST *IFSTAT, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { // IF STATEMENT
	var ok bool
	var error string

	if AST.ifexp != nil {
		ok, error, _, _ = EXPR(AST.ifexp, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	if AST.statseq != nil {
		ok, error = SQ(AST.statseq, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	if AST.elseifexp != nil {
		ok, error, _, _ = EXPR(AST.elseifexp, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	if AST.eistatseq != nil {
		ok, error = SQ(AST.eistatseq, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	if AST.estatseq != nil {
		ok, error = SQ(AST.estatseq, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	return true, "OK"
}

func WT(AST *WSTAT, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //WHILE STATEMENT
	if AST.exp != nil {
		ok, error, _, _ := EXPR(AST.exp, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	if AST.statseq != nil {
		ok1, error1 := SQ(AST.statseq, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	return true, "OK"

}

func FT(AST *FSTAT, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //FOR STATEMENT
	if AST.id != nil {
		ok, error, name, lno := ID(AST.id, m, n, scope, lex)
		if ok == true {
			ok0, error0 := TableCheck(name, lno, m, n, 1, lex)
			if ok0 == false {
				return ok0, error0
			}
		} else {
			return ok, error
		}
	}
	if AST.exp1 != nil {
		ok1, error1, _, _ := EXPR(AST.exp1, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	if AST.exp2 != nil {
		ok2, error2, _, _ := EXPR(AST.exp2, m, n, scope, lex)
		if ok2 == false {
			return ok2, error2
		}
	}
	if AST != nil {

	}
	if AST.statseq != nil {
		ok3, error3 := SQ(AST.statseq, m, n, scope, lex)
		if ok3 == false {
			return ok3, error3
		}
	}

	return true, "OK"
}

func EXPR(AST *EXP, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //EXPRESSION
	var value, value3, error1 string
	var ok1 bool
	var lno int

	if AST.se != nil {
		ok1, error1, value, lno = SExp(AST.se, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1, value, lno
		}
	}
	if AST.re != nil {
		REOP(AST.re, m, n, scope, lex)

	}
	if AST.se2 != nil {
		ok1, error1, value3, lno = SExp(AST.se2, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1, value3, lno
		}
	}

	return true, "OK", value, lno
}

func SExp(AST *SE, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //simple expression
	var value, error1 string
	var ok1 bool
	var lno int

	if AST != nil {

	}
	if AST.term != nil {
		ok1, error1, value, lno = TM(AST.term, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1, value, lno
		}
	}
	if AST.addop != nil {
		value = AOP(AST.addop, m, n, scope, lex)
	}
	if AST.term2 != nil {
		ok1, error1, value, lno = TM(AST.term2, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1, value, lno
		}
	}
	return true, "OK", value, lno
}

func REOP(AST *RE, m map[int]M, n map[int]N, scope int, lex *LexType) { //Relation Operator
	if AST != nil {
		//return AST.reop
	}
	//return ""
}

func AOP(AST *ADDOP, m map[int]M, n map[int]N, scope int, lex *LexType) string { //ADD Operator
	if AST != nil {
		return AST.op
	}
	return ""
}

func MUP(AST *MULOP, m map[int]M, n map[int]N, scope int, lex *LexType) string { //MUL Operator
	if AST != nil {
		return AST.op
	}
	return ""

}

func FCT(AST *FACT, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //factor
	var value, error string
	var lno int
	var ok bool
	if AST.integer != nil {
		ok, error, value, lno := ID(AST.integer, m, n, scope, lex)
		return ok, error, value, lno
	}
	if AST.str != nil {
		ok, error, value, lno := ID(AST.str, m, n, scope, lex)
		return ok, error, value, lno
	}
	if AST.bol1 != nil {
		ok, error, value, lno := ID(AST.bol1, m, n, scope, lex)
		return ok, error, value, lno
	} // String data
	if AST.bol2 != nil {
		ok, error, value, lno := ID(AST.bol2, m, n, scope, lex)
		return ok, error, value, lno
	}
	if AST.dsg != nil {
		ok, error, value, lno := DG(AST.dsg, m, n, scope, lex)
		if ok == false {
			return ok, error, value, lno
		}
	}
	if AST.ap != nil {
		APA(AST.ap, m, n, scope, lex)
	}
	if AST.exp != nil {
		ok, error, value, lno = EXPR(AST.exp, m, n, scope, lex)
		if ok == false {
			return ok, error, value, lno
		}
	}
	if AST.til != nil {
		ok, error, value, lno := ID(AST.til, m, n, scope, lex)
		return ok, error, value, lno
	}
	if AST.fac != nil {
		ok, error, value, lno = FCT(AST.fac, m, n, scope, lex)
		if ok == false {
			return ok, error, value, lno
		}
	}
	return true, "OK", value, lno

}

func TM(AST *TERM, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //term
	var value, error1 string
	var lno int
	var ok1 bool
	if AST.fact != nil {
		ok1, error1, value, lno = FCT(AST.fact, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1, value, lno
		}
	}
	if AST.mulop != nil {
		value = MUP(AST.mulop, m, n, scope, lex)
	}
	if AST.fact2 != nil {
		ok1, error1, value, lno = FCT(AST.fact2, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1, value, lno
		}
	}
	return true, "OK", value, lno
}

func APA(AST *AP, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //actual parameters
	if AST.expl != nil {
		ok, error := EXPL(AST.expl, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	return true, "OK"
}

func EXPL(AST *EXPLIST, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) {

	if AST.e != nil {
		for i := 0; i < len(AST.e); i++ {
			if AST.e[i] != nil {
				ok, error, _, _ := EXPR(AST.e[i], m, n, scope, lex)
				if ok == false {
					return false, error
				} else {
					return true, "OK"
				}

			}
		}
	}
	return true, "OK"
}

func DG(AST *DESG, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //designator
	var name, error string
	var lno int
	var ok bool
	if AST.id != nil {
		ok, error, name, lno = ID(AST.id, m, n, scope, lex)
		if ok == false {
			return ok, error, name, lno
		}
		ok1, error1 := TableCheck(name, lno, m, n, 0, lex)
		if ok1 == false {
			return ok1, error1, name, lno
		}
	}

	if AST.sel != nil {
		ok2, error2, name2, lno2 := SETR(AST.sel, m, n, scope, lex)
		if ok2 == false {
			return ok2, error2, name2, lno2
		}
	}

	return true, "ok", name, lno
}

func SETR(AST *SEL, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { // SELECTOR
	var name, error string
	var lno int
	var ok bool

	if AST.exp != nil {
		ok, error, name, lno = EXPR(AST.exp, m, n, scope, lex)
		if ok == false {
			return ok, error, name, lno
		}
	}
	return true, "OK", name, lno
}

func TD(AST *TDEC, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //TYPE DECLARATION
	var error1, dtype, name string
	var ok1 bool
	var lno int
	if AST.id != nil {
		ok1, error1, name, lno = ID(AST.id, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	if AST.at != nil {
		ok1, error1, dtype, _ = ART(AST.at, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	if scope == 0 {
		ok1, error1 := SymTable(name, dtype, lno, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	} else if scope == 1 {
		ok2, error2 := PsymTable(name, dtype, lno, m, n, scope, lex)
		if ok2 == false {
			return ok2, error2
		}
	}

	return true, "OK"
}

func ID(AST *IDENT, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //IDENT

	if AST != nil {
		return true, "OK", AST.id, AST.lineno
	}

	return false, "ERROR", AST.id, AST.lineno
}

func ART(AST *AT, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //	ARRAY TYPE

	var lno int
	var dtype string
	if AST != nil {

	}
	if AST.ty != nil {
		dtype := TE(AST.ty, m, n, scope, lex)
		return true, "OK", dtype, lno
	}
	return true, "OK", dtype, lno
}

func TE(AST *TYPE, m map[int]M, n map[int]N, scope int, lex *LexType) string { //  TYPE
	if AST != nil {
		return AST.id.id
	}
	return ""
}

func VD(AST *VDEC, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //VARIABLE DECLARATION
	name := []string{}
	var lno int
	var dtype, error1 string
	var ok1 bool
	if AST.Il != nil {
		ok1, error1, name, lno = IL(AST.Il, m, n, scope, lex)
		if ok1 == false {
			return ok1, error1
		}
	}
	if AST.ty != nil {
		dtype = TE(AST.ty, m, n, scope, lex)

	}
	for j := 0; j < len(name); j++ {
		if scope == 0 {
			ok1, error1 := SymTable(name[j], dtype, lno, m, n, scope, lex)
			if ok1 == false {
				return ok1, error1
			}
		} else if scope == 1 {
			ok2, error2 := PsymTable(name[j], dtype, lno, m, n, scope, lex)
			if ok2 == false {
				return ok2, error2
			}
		}
	}

	return true, "OK"
}

func IL(AST *IDENTLIST, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, []string, int) { // IDENT LIST
	name := []string{}
	var lno, lno2 int
	var ok bool
	var error, nam string
	for i := 0; i < len(AST.idl); i++ {
		if AST.idl[i] != nil {
			ok, error, nam, lno2 = ID(AST.idl[i], m, n, scope, lex)
			name = append(name, nam)
			lno = lno2
		}
	}
	return ok, error, name, lno
}

func PD(AST *PDEC, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //procedure DECLARATION

	if AST.phead != nil {
		ok, error := PH(AST.phead, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	if AST.pbody != nil {
		ok2, error2 := PB(AST.pbody, m, n, scope, lex)
		if ok2 == false {
			return ok2, error2
		}
	}

	if AST.id != nil {
		_, _, name, lno := ID(AST.id, m, n, scope, lex)

		if name != AST.phead.id.id {
			msg := "IDENT MISMATCH at LN " + strconv.Itoa(lno)
			return false, msg
		}

	}
	return true, "OK"

}

func PH(AST *PHEAD, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //PROCEDURE HEAD

	if AST.id != nil {
		_, _, name, lno := ID(AST.id, m, n, scope, lex)
		dtype := "CHAR"
		ok1, error1 := SymTable(name, dtype, lno, m, n, 0, lex)
		if ok1 == false {
			return ok1, error1
		}

	}

	for i := 0; i < len(AST.fpsec); i++ {
		if AST.fpsec[i] != nil {
			ok2, error2 := FSEC(AST.fpsec[i], m, n, 1, lex)
			if ok2 == false {
				return ok2, error2
			}
		}
	}
	if AST.id2 != nil {
		_, _, name2, ln2 := ID(AST.id2, m, n, 1, lex)

		ok3, error3 := TableCheck(name2, ln2, m, n, 1, lex)
		if ok3 == false {
			return ok3, error3
		}
		//fmt.Println(name2)
	}
	return true, "OK"
}

func PB(AST *PBODY, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //PROCEDURE BODY
	if AST.decseq != nil {
		ok, error := DQ(AST.decseq, m, n, 1, lex)
		if ok == false {
			return ok, error
		}
	}
	if AST.statseq != nil {
		ok2, error2 := SQ(AST.statseq, m, n, 1, lex)
		if ok2 == false {
			return ok2, error2
		}
	}
	if AST.exp != nil {
		ok3, error3, _, _ := EXPR(AST.exp, m, n, 1, lex)
		if ok3 == false {
			return ok3, error3
		}
	}
	return true, "OK"
}

func FSEC(AST *FPSEC, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //FPSECTION
	var name, typ, error1, error2 string
	var lno int
	var ok1, ok2 bool

	for i := 0; i < len(AST.id); i++ {
		if AST.id[i] != nil {
			ok1, error1, name, lno = ID(AST.id[i], m, n, scope, lex)

			if ok1 == false {
				return ok1, error1
			}

		}
	}
	if AST.ftype != nil {
		ok2, error2, typ, _ = FTY(AST.ftype, m, n, scope, lex)
		if ok1 == false {
			return ok1, error2
		}
	}

	if ok1 && ok2 {

		ok, error := PsymTable(name, typ, lno, m, n, scope, lex)
		if ok == false {
			return ok, error
		}
	}
	return true, "OK"

}

func FTY(AST *FTYPE, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string, string, int) { //Formal Type

	var name, error string
	var ok bool
	var lno int

	if AST.id != nil {
		ok, error, name, lno = ID(AST.id, m, n, scope, lex)
		if ok == false {
			return ok, error, name, lno
		}

	}
	return true, "OK", name, lno
}

var n map[int]N

type N struct {
	dname string
	dtype string
}

func PsymTable(name string, dtype string, lno int, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { // PROCEDURE SYMBOL TABLE
	count := 0
	count2 := 0
	count3 := 0

	if dtype != "INTEGER" && dtype != "CHAR" && dtype != "BOOLEAN" {
		temp := dtype
		for i := 0; i < len(n); i++ {
			if n[i].dname == dtype {
				if n[i].dtype != "INTEGER" && dtype != "CHAR" && dtype != "BOOLEAN" {
					dtype = n[i].dname
					continue
				}
				if n[i].dtype == "INTEGER" && dtype == "CHAR" && dtype == "BOOLEAN" {
					break
				}
			} else {
				count2++
			}
		}
		if count2 == len(n) {
			for j := 0; j < len(m); j++ {
				if m[j].dname == temp {
					if m[j].dtype != "INTEGER" && dtype != "CHAR" && dtype != "BOOLEAN" {
						temp = m[j].dname
						continue
					}
					if m[j].dtype == "INTEGER" && dtype == "CHAR" && dtype == "BOOLEAN" {
						break
					}

				} else {
					count3++
				}
				if count3 == len(m) {
					return false, lex.ErrorMessage("INVALID DATATYPE " + name + ", " + strconv.Itoa(lno))
				}

			}
		}

	}

	for i := 0; i < len(n); i++ {
		if name != n[i].dname {
			count++
		}
	}
	if count == len(n) {
		n[index2] = N{name, dtype}
		index2++
	} else {
		return false, lex.ErrorMessage("VAR REDECLARED " + name + ", " + strconv.Itoa(lno))

	}

	return true, "OK"

}

func PTableCheck(name string, lno int, n map[int]N, scope int, lex *LexType) (bool, string) { //CHECK SYMBOL TABLE
	count := 0

	for i := 0; i < len(m); i++ {
		if name != n[i].dname {
			count++
		}
	}
	if count == len(n) {
		return false, lex.ErrorMessage("VAR NOT DECLARED " + name + ", " + strconv.Itoa(lno))

	}

	return true, "OK"
}

var m map[int]M

type M struct {
	dname string
	dtype string
	//line int
}

func SymTable(name string, dtype string, lno int, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //Creating symbol table
	count := 0
	count2 := 0
	i := 0
	// CHECK for valid DATAtype
	if dtype != "INTEGER" && dtype != "CHAR" && dtype != "BOOLEAN" {
		for i = 0; i < len(m); i++ {
			if m[i].dname == dtype {
				if m[i].dtype != "INTEGER" && dtype != "CHAR" && dtype != "BOOLEAN" {
					dtype = m[i].dname
					continue
				}
				if m[i].dtype == "INTEGER" && dtype == "CHAR" && dtype == "BOOLEAN" {
					break
				}
			} else {
				count2++
			}
		}
		if count2 == len(m) {
			return false, lex.ErrorMessage("INVALID DATATYPE " + name + ", " + strconv.Itoa(lno))
		}

	}

	if scope == 0 { // MODULE SCOPE
		for i := 0; i < len(m); i++ {
			if name != m[i].dname {
				count++
			}
		}
		if count == len(m) {
			m[index] = M{name, dtype}
			index++
		} else {
			return false, lex.ErrorMessage("VAR REDECLARED " + name + ", " + strconv.Itoa(lno))
		}
	}

	if scope == 1 { // PROCEDURE SCOPE
		for i := 0; i < len(n); i++ {
			if name != n[i].dname {
				count++
			}
		}

		if count == len(n) {
			n[index2] = N{name, dtype}
			index2++
		} else {
			return false, lex.ErrorMessage("VAR REDECLARED " + name + ", " + strconv.Itoa(lno))
		}
	}

	return true, "OK"

}

func TableCheck(name string, lno int, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) { //CHECK SYMBOL TABLE
	count, count1, count2 := 0, 0, 0

	if scope == 0 {
		for i := 0; i < len(m); i++ {
			if name != m[i].dname {
				count++
			}
		}
		if count == len(m) {
			return false, lex.ErrorMessage("VAR NOT DECLARED " + name + ", " + strconv.Itoa(lno))
		}
	}
	if scope == 1 {
		for i := 0; i < len(n); i++ {
			if name != n[i].dname {
				count1++
			}
		}
		for j := 0; j < len(m); j++ {
			if name != m[j].dname {
				count2++
			}
		}
		if count1 == len(n) && count2 == len(m) {
			return false, lex.ErrorMessage("VAR NOT DECLARED " + name + ", " + strconv.Itoa(lno))
		}
	}
	return true, "OK"
}

func Checksemantics(name string, value string, lno int, m map[int]M, n map[int]N, scope int, lex *LexType) (bool, string) {
	var j int
	var temp string

	if value == "OK" {
	}
	// CHECK FOR INTEGERS
	_, msg := strconv.Atoi(value)
	if msg == nil {
		for i := 0; i < len(m); i++ {
			if name == m[i].dname {
				if m[i].dtype == "INTEGER" {

					return true, "OK"
				} else if m[i].dtype != "INTEGER" {
					temp = m[i].dtype
					for j = 0; j < len(m); j++ {
						if m[j].dname == temp && m[j].dtype != "INTEGER" {
							temp = m[j].dtype
							continue
						} else if m[j].dname == temp && m[j].dtype == "INTEGER" {

							return true, "OK"
						}
					}

					if j == len(m) {
						return false, lex.ErrorMessage("TYPE MISMATCH " + name + ", " + strconv.Itoa(lno))
					}
				}
			}
		} // CHECK FOR BOOLEAN
	} else if value == "FALSE" || value == "TRUE" {
		for i := 0; i < len(m); i++ {
			if name == m[i].dname {
				if m[i].dtype != "BOOLEAN" {
					return false, lex.ErrorMessage("TYPE MISMATCH " + name + ", " + strconv.Itoa(lno))
				}
			}
		}
	} else { //CHECK FOR CHAR(STRING)

		for i := 0; i < len(m); i++ {
			if name == m[i].dname {
				if m[i].dtype == "CHAR" {

					return true, "OK"
				} else if m[i].dtype != "CHAR" {
					temp = m[i].dtype
					for j = 0; j < len(m); j++ {
						if m[j].dname == temp && m[j].dtype != "CHAR" {
							temp = m[j].dtype
							continue
						} else if m[j].dname == temp && m[j].dtype == "CHAR" {

							return true, "OK"
						}
					}
				}

				if j == len(m) {
					return false, lex.ErrorMessage("TYPE MISMATCH " + name + ", " + strconv.Itoa(lno))
				}
			}
		}
	}

	return true, "OK"

}
