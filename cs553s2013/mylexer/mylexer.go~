/* Lexer in Go for Oberon Language */
/* AUTHOR - KARAN CHADHA ; DATE-1/2013 ; CS553 Spring13 */

package mylexer

import (
	"fmt"
	"github.com/proebsting/cs553s2013/lexer"
	"unicode"
)

var book string
var booklength int
var bookmark int
var typeofword int

type wordtype int

const (
	Identifier = 0
	Operator   = 1
	Keyword    = 2
	String     = 3
	Number     = 4
	Eof        = 5
	Error      = 6
)

type TestToken struct { // Struct for Storing Token Values
	enum     int 
	value    string
	location int
	line     int
	column   int
}

func (x *TestToken) Column() int {
	return x.column
}
func (x *TestToken) Line() int {
	return x.line
}
func (x *TestToken) Value() string {
	return x.value
}
func (x *TestToken) Enum() int {
	return x.enum
}
func (x *TestToken) Location() int {
	return 0
}

func DumpToken(t lexer.Token) (s string) { // Function to Dump Token Info
	//s = fmt.Sprintf("Token{ Enum=%s(%d), Value=%s, Line=%d, Column=%d, Location=%d }", lexer.Enum2String(t.Enum()), t.Enum(), t.Value(), t.Line(), t.Column(), t.Location())
	return 
}

func Wordtype() int { // Function to Return Token Type
	return typeofword
}

func Check() int { // Check for not reading Beyond the file
	if bookmark == booklength {
		typeofword = Eof
		return -1
	}
	return 0
}

var t = TestToken{value: "", enum: 0, location: 0, line: 1, column: 1}

func Lexer(input string, output chan lexer.Token) { // Reading the Input as a Book , with Bookmark as the current charachter
	//output := make(chan lexer.Token, 100)
	output <- &t
	//fmt.Println(len(output))
	
	book = input
	booklength = len(book)
	bookmark = 0
	t.line = 0
	t.column = 1
	t.enum = 0
	t.value = ""
	
	//fmt.Println(bookmark)
	Getnextword()
	
	for Wordtype() <= 4 { //Dumping the Token on the Display and Channel
		fmt.Printf("%v",DumpToken(&t))
		output <- &t
		t.column = t.column + len(t.value)
		Getnextword()
	}
	if Wordtype() == 5 { // EOF

		t.enum = lexer.EOF
		t.value = "<EOF>"
		fmt.Printf("%v",DumpToken(&t))
		

		return
	} else if Wordtype() == 6 { //ERROR
		t.enum = lexer.ERROR
		fmt.Println(DumpToken(&t))
		output <- &t
		return
	}
	
}

func Getnextword() {
	t.value = ""
	t.enum = 0
	typeofword = Error
	value := Check()
	if value == -1 {
		return
	}
	Skipwhitespace()
	value3 := Check()
	if value3 == -1 {
		return
	}
	if book[bookmark] == '(' && book[bookmark+1] == '*' { //Check for Comment
		result := Skipcomment()
		if result == -1 {
			return
		}
	}
	Skipwhitespace()
	if unicode.IsLetter(rune(book[bookmark])) { //check for Identifier and Keyword
		Getidentifier()
		return
	}
	value, length, valid := lexer.Operator(string(book[bookmark])) //check for Operator
	if valid && length >= 1 {
		Getoperator(value)
		return
	}
	if book[bookmark] == '"' { // Check for string
		Getstring()
		return
	}
	if unicode.IsDigit(rune(book[bookmark])) { // Check for Number
		Getnumber()
		return
	}
	value2 := Check()
	if value2 == -1 {
		return
	}

}

func Skipcomment() int { //Ignoring Comment
	bookmark += 2
	ocolumn := t.column + 2
	comment := 1
	oline := t.line
	value := Check()
	if value == -1 {
		return -1
	}
	for comment != 0 && bookmark < booklength-1 {
		if bookmark == booklength-1 {
			typeofword = Error
			t.column = 1
			return -1
		}
		if book[bookmark] == '(' && book[bookmark+1] == '*' { //Starting of Comment
			comment++
			bookmark += 2
			ocolumn += 2
		} else if book[bookmark] == '*' && book[bookmark+1] == ')' { //Ending of Comment
			comment--
			bookmark += 2
			ocolumn += 2
			t.column = ocolumn
			t.line = oline
			if comment == 0 {
				return 0
			}
		} else if book[bookmark] == '\n' { //Newline within Comment
			oline++
			bookmark++
			ocolumn = 1
		} else { //Ignore the Rest
			bookmark++
			ocolumn++
		}

	}
	typeofword = Error
	return -1
}

func Skipwhitespace() {
	value := Check()
	if value == -1 {
		return
	}
	for book[bookmark] == ' ' || book[bookmark] == '\n' || book[bookmark] == '\t' || book[bookmark] == '\r' {

		if book[bookmark] == '\n' {
			t.line++
			bookmark++
			t.column = 1
			value2 := Check()
			if value2 == -1 {
				return
			}

		} else {
			bookmark++
			t.column++
		}
	}
}

func Getidentifier() {
	value := Check()
	if value == -1 {
		return
	}
	for unicode.IsLetter(rune(book[bookmark])) || unicode.IsDigit(rune(book[bookmark])) {
		t.value = t.value + string(book[bookmark])
		bookmark++
	}
	token, match := lexer.Keyword(t.value)
	if match && token != -1 { // If Keyword found Return
		typeofword = Keyword
		t.enum = token
	} else { //Else Return Identifier
		typeofword = Identifier
		t.enum = lexer.IDENT
	}
}

func Getoperator(value int) {
	value2 := Check()
	if value2 == -1 {
		return
	}
	t.value = string(book[bookmark])
	t.enum = value
	if bookmark+1 != booklength {
		value2, length2, valid2 := lexer.Operator(string(book[bookmark+1]))
		if value2 != -1 && valid2 && length2 >= 1 {
			input := string(book[bookmark]) + string(book[bookmark+1])
			value3, length3, valid3 := lexer.Operator(input)
			if valid3 && length3 > 1 {
				t.value = input
				bookmark += 2
				t.enum = value3
				typeofword = Operator
				return
			}
		}
	}
	bookmark++
	typeofword = Operator
}

func Getstring() {
	value2 := Check()
	if value2 == -1 {
		return
	}
	bookmark++
	for book[bookmark] != '"' {

		t.value = t.value + string(book[bookmark])
		bookmark++
		if bookmark > booklength || book[bookmark] == '\n' {
			t.value = "String-Error"
			typeofword = Error
			return
		}
	}
	bookmark++
	t.value = "\"" + t.value + "\""
	typeofword = String
	t.enum = lexer.STRING
}

func Getnumber() {
	value2 := Check()
	if value2 == -1 {
		return
	}
	var reachedDOT = false
	var reachedED = false
	for unicode.IsDigit(rune(book[bookmark])) { //Start With a Digit

		t.value = t.value + string(book[bookmark])
		bookmark++
		value2 := Check()
		if value2 == -1 {
			return
		}
		for (book[bookmark] == 'A' || book[bookmark] == 'B' || book[bookmark] == 'C' || book[bookmark] == 'D' || book[bookmark] == 'E' || book[bookmark] == 'F') && reachedDOT == false { //Hexa
			t.value = t.value + string(book[bookmark])
			bookmark++
		}
		if book[bookmark] == 'H' { //End with H (INTEGER)
			t.value = t.value + string(book[bookmark])
			bookmark++
			t.enum = lexer.INTEGER
			typeofword = Number
		}
		if book[bookmark] == 'X' { //End with X (STRING)
			t.value = t.value + string(book[bookmark])
			bookmark++
			t.enum = lexer.STRING
			typeofword = String
			return
		}
		if book[bookmark] == '.' && reachedDOT == false { //Decimal "." (REAL)
			t.value = t.value + string(book[bookmark])
			bookmark++
			reachedDOT = true
		}
		if book[bookmark] == 'E' || book[bookmark] == 'D' { //Base Multiplier (REAL)
			t.value = t.value + string(book[bookmark])
			bookmark++
			reachedED = true
		}
		if book[bookmark] == '+' || book[bookmark] == '-' {

			if reachedED {
				t.value = t.value + string(book[bookmark])
				bookmark++
			}
		}
	}
	if (reachedED || reachedDOT) && unicode.IsDigit(rune(t.value[len(t.value)-1])) { //Conditions For Identifying Integer or Real else Error
		t.enum = lexer.REAL
		typeofword = Number
	} else if reachedED == false && reachedDOT == false {
		t.enum = lexer.INTEGER
		typeofword = Number
	} else if reachedED == true && reachedDOT == false {
		t.enum = lexer.ERROR
		typeofword = Error
	} else if reachedDOT && unicode.IsDigit(rune(book[bookmark])) {
		t.enum = lexer.REAL
		typeofword = Number
	} else {
		t.enum = lexer.ERROR
		typeofword = Error
		return
	}
}
