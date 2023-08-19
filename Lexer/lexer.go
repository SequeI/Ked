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
		l.char = byte(EOF)
	}
}

func isWhitespace(char byte) bool {
	// FUTURE COLUMN + ROW TRACKER FOR ERROR ACCURACY
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

// tokenizer function
func Tokenizer(input string) string {
	lexer := NewLexer(input)

	for lexer.char != byte(EOF) {
		lexer.nextChar()
		//lint:ignore SA4017 Ignore this warning because the isWhitespace function is used for code organization.
		if isWhitespace(lexer.char) {
			// FUTURE COLUMN + ROW TRACKER FOR ERROR ACCURACY
			continue
		}
	}

	return " "
}
