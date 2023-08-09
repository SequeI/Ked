package lexer

type TokenType int

const (
	// End of file
	EndOfFile TokenType = iota

	// Identifier & Literal
	Identifier
	Number
	Float

	//Operators
	Add
	Minus
	Assign
	Mul
	Div

	//Delimiters
	EndOfLine
	Lparen
	Rparen

	//Keywords
	Function
	Let
)

type Token struct {
	Type  TokenType
	Value string
}
