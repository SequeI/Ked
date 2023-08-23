package lexer

import (
	"strconv"
)

// FUNCTIONS & METHODS

// Character checking involves evaluating individual characters to determine their properties or significance within the code.
// This process helps identify keywords, whitespace, symbols, and other elements, enabling proper tokenization and further analysis.
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

// The Tokenizer function processes the input code by breaking it down into individual tokens.
// It iterates through the characters, identifying keywords, identifiers, and numbers,
// and creates corresponding tokens for each recognized element.
func Tokenizer(input string) []Token {

	current := 0
	tokens := []Token{}
	// Iterating through individual characters and performing specific operations based on their values.
	for current < len(input) {

		char := input[current]
		// This section checks if the character corresponds to a valid token type in the hashmap,
		// such as whether it is a closing parenthesis or another recognized symbol.
		if _, ok := keywordTokens[string(char)]; ok && (string([]rune(input)[current+1]) == " " || current == len(input)-1) {
			tokens = append(tokens, Token{Type: keywordTokens[string(char)], Value: string(char)})
			current += 2
			continue
		}
		// Whitespace checker examines characters to identify whether they are spaces, assisting in seperating and formatting the source code.
		if char == ' ' {
			current++
			continue
		}
		// This section handles the processing of characters that are either letters or digits. It accumulates these characters into a value string,
		// effectively constructing complete identifiers or integer literals. The loop continues until characters remain that are either letters or digits.
		// If the accumulated value is numeric, an INT token is created; otherwise, an IDENTIFIER token is generated.
		if isNumber(string(char)) || isLetter(string(char)) {
			value := ""

			for isNumber(string(char)) || isLetter(string(char)) {
				value += string(char)
				current++
				if current < len(input) {
					char = input[current]
				} else {
					break
				}
			}

			if isNumeric(value) {
				tokens = append(tokens, Token{Type: INT, Value: value})
			} else {
				tokens = append(tokens, Token{Type: IDENTIFIER, Value: value})
			}
			continue
		}
		break
	}
	// Once all characters have been processed and tokens constructed, the function returns a slice
	// containing each individual token that has been identified during the tokenization process.
	return tokens
}
