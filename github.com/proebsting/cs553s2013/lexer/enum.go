package lexer

const (
	ERROR        = -1
	EOF          = iota
	IDENT        // identifiers
	STRING       // string literals
	HASH         // #
	AMP          // &
	LPAREN       // (
	RPAREN       // )
	STAR         // *
	PLUS         // +
	COMMA        // ,
	MINUS        // -
	DOT          // .
	DOTDOT       // ..
	SLASH        // /
	COLON        // :
	COLONEQUAL   // :=
	SEMICOLON    // ;
	LESS         // <
	LESSEQUAL    // <=
	EQUAL        // =
	GREATER      // >
	GREATEREQUAL // >=
	LSQUARE      // [
	RSQUARE      // ]
	CARAT        // ^
	LCURLY       // {
	BAR          // |
	RCURLY       // }
	TILDE        // ~
	ARRAY
	BEGIN
	BY
	CASE
	CONST
	DIV
	DO
	ELSE
	ELSIF
	END
	FALSE
	FOR
	IF
	IMPORT
	IN
	INTEGER
	IS
	MOD
	MODULE
	NIL
	OF
	OR
	POINTER
	PROCEDURE
	REAL
	RECORD
	REPEAT
	RETURN
	THEN
	TO
	TRUE
	TYPE
	UNTIL
	VAR
	WHILE
)
