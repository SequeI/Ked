package lexer

type Lexer struct {
	input  string
	pos    int
	char   byte
	length int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		pos:    -1,
		length: len(input),
		char:   0,
	}
}

// tokenizer function
func Tokenizer(input string) string {
	lexer := NewLexer(input)

	return " "
}
