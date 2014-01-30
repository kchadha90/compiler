package lexer

type Token interface {
	Value() string
	Location() int
	Line() int
	Column() int
	Enum() int
}
