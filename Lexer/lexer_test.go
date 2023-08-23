package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
	remember x = 4 like
	remember y = 2 like

	eh ( jaKnow ( x ) ) {
		remember z = 1 like
	}
	orEh ( jaKnow ( y ) ) {
		remember z = 2 like
	}
	orEvenJust {
		remember z = 0 like
	}
	7 awayFrom 4
	4 - 1 
	`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LET, "remember"},
		{IDENTIFIER, "x"},
		{ASSIGN, "="},
		{INT, "4"},
		{EOL, "like"},
		{LET, "remember"},
		{IDENTIFIER, "y"},
		{ASSIGN, "="},
		{INT, "2"},
		{EOL, "like"},
		{IF, "eh"},
		{LPAREN, "("},
		{IDENTIFIER, "jaKnow"},
		{LPAREN, "("},
		{IDENTIFIER, "x"},
		{RPAREN, ")"},
		{RPAREN, ")"},
		{LBRACKET, "{"},
		{LET, "remember"},
		{IDENTIFIER, "z"},
		{ASSIGN, "="},
		{INT, "1"},
		{EOL, "like"},
		{RBRACKET, "}"},
		{ELIF, "orEh"},
		{LPAREN, "("},
		{IDENTIFIER, "jaKnow"},
		{LPAREN, "("},
		{IDENTIFIER, "y"},
		{RPAREN, ")"},
		{RPAREN, ")"},
		{LBRACKET, "{"},
		{LET, "remember"},
		{IDENTIFIER, "z"},
		{ASSIGN, "="},
		{INT, "2"},
		{EOL, "like"},
		{RBRACKET, "}"},
		{ELSE, "orEvenJust"},
		{LBRACKET, "{"},
		{LET, "remember"},
		{IDENTIFIER, "z"},
		{ASSIGN, "="},
		{INT, "0"},
		{EOL, "like"},
		{RBRACKET, "}"},
		{INT, "7"},
		{MINUS, "awayFrom"},
		{INT, "4"},
		{MINUS, "-"},
		{INT, "1"},
	}

	tokens := Tokenizer(input)
	fmt.Println(tokens)

	if len(tokens) != len(tests) {
		t.Fatalf("Mismatched number of tokens: expected %d, got %d", len(tests), len(tokens))
	}

	for i, token := range tests {
		if tokens[i].Type != token.expectedType {
			t.Fatalf("Test %d: Expected type %q, got %q", i, token.expectedType, tokens[i].Type)
		}
		if tokens[i].Value != token.expectedLiteral {
			t.Fatalf("Test %d: Expected literal %q, got %q", i, token.expectedLiteral, tokens[i].Value)
		}
	}
}
