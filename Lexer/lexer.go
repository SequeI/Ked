package lexer

import (
	"unicode"
)

type Lexer struct {
	input  string
	pos    int
	length int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		pos:    0,
		length: len(input),
	}
}

// Functions + Methods
func isWhitespace(char rune) bool {
	return unicode.IsSpace(char)
}

// tokenizer function
func Tokenizer(input string) string {

	tokens := []Token{}
	lexer := NewLexer(input)

	for lexer.pos < lexer.length {
		char := lexer.input[lexer.pos]

		if isWhitespace(rune(char)) {
			lexer.pos++
			continue
		}
		tempV := string(char)

		for lexer.pos+1 < lexer.length {
			nextChar := lexer.input[lexer.pos+1]
			if isWhitespace(rune(nextChar)) {
				break
			}
			tempV += string(nextChar)
			lexer.pos++
		}
		if tokenValue, ok := keywordTokens[tempV]; ok {
			token := Token{Type: tokenType, Value: tokenValue}
			tokens = append(tokens, token)
		} else {
			
		}
	

	return " "
}
