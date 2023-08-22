package lexer

// tokenizer function
func Tokenizer(input string) []Token {

	input += "\n"
	current := 0
	tokens := []Token{}

	for current < len([]rune(input)) {

		char := string([]rune(input)[current])

		if _, ok := keywordTokens[char]; ok && string([]rune(input)[current+1]) == " " {
			tokens = append(tokens, Token{Type: keywordTokens[char], Value: char})
		}
	}
	return tokens
}
