package lexer

type TokenType int

const (
	// End of file
	EOF TokenType = iota

	// Identifier & Literal
	Identifier
	Number

	//Operators
	Add
	Minus
	Assign

	//Delimiters
	EndOfLine

	//Keywords
	Function
	Let
)

type Token struct {
	Type  TokenType
	Value string
}
