package lexer

import (
	"strconv"
)

// Functions & Methods
func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z') {
		return true
	}
	return false
}

func isNumber(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= '0' && n <= '9' {
		return true
	}
	return false
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// tokenizer function
func Tokenizer(input string) []Token {

	input += "\n"
	current := 0
	tokens := []Token{}

	for current < len([]rune(input)) {

		char := string([]rune(input)[current])

		if _, ok := keywordTokens[char]; ok && (string([]rune(input)[current+1]) == " " || current == len(input)-1) {
			tokens = append(tokens, Token{Type: keywordTokens[char], Value: char})
			current += 2
			continue
		}
		if char == " " {
			current++
			continue
		}
		if isNumber(char) || isLetter(char) {
			value := ""

			for isNumber(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}
			for isLetter(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}
			if isNumeric(value) {
				tokens = append(tokens, Token{Type: INT, Value: value})
				continue
			} else {
				tokens = append(tokens, Token{Type: IDENTIFIER, Value: value})
				continue
			}
		}
		break
	}
	return tokens
}
