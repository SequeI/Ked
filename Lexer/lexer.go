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

func isAlphabetChar(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

// tokenizer function
func Tokenizer(input string) []Token {

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
		if _, ok := keywordTokens[tempV]; ok {
			token := Token{Type: keywordTokens[tempV], Value: tempV}
			tokens = append(tokens, token)
		} else {
			if isAlphabetChar(tempV) {
				token := Token{Type: IDENTIFIER, Value: tempV}
				tokens = append(tokens, token)
			} else {
				// Probably integer if doesn't contain letters lol, will do better checks later
				token := Token{Type: INT, Value: tempV}
				tokens = append(tokens, token)
			}
		}
	}

	// Add EOF token at the end
	token := Token{Type: EOF, Value: "EndOfFile"}
	tokens = append(tokens, token)

	return tokens
}
