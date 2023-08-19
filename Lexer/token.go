package lexer

type TokenType int

const (
	// End of file
	EOF TokenType = iota

	// Keywords
	LET
	IF
	ELIF
	ELSE
	WHILE

	// Identifier
	IDENTIFIER

	//Literals
	INT

	// Operators
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO
	ASSIGN

	// Delimiters and Symbols
	LPAREN
	RPAREN
)

type Token struct {
	Type  TokenType
	Value string
}

// Hash Map
var keywordTokens = map[string]TokenType{
	"remember":    LET,
	"eh":          IF,
	"orEh":        ELIF,
	"orEvenJust":  ELSE,
	"eraGoOnSure": WHILE,
	"plus":        PLUS,
	"awayFrom":    MINUS,
	"times":       MULTIPLY,
	"into":        DIVIDE,
	"%":           MODULO,
	"=":           ASSIGN,
	"(":           LPAREN,
	")":           RPAREN,
}
