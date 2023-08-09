package lexer

// import "unicode"

// Inserting source code into a struct

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

// Methods for tokenizing source code

func (l *Lexer) nextChar() {
	l.pos++
	if l.pos < l.length {
		l.char = l.input[l.pos]
	} else {
		l.char = byte(EndOfFile)
	}
}

// tokenizer function
func Tokenizer(input string) string {
	lexer := NewLexer(input)
	for lexer.char != byte(EndOfFile) {
		lexer.nextChar()
	}

	return " "
}
