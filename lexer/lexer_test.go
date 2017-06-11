package lexer

import (
	"monkeyinterpreter/token"
	"testing"
)

// TestNextToken tests lexer.NextToken()
func TestNextToken(assert *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for index, tokenType := range tests {
		nextToken := lexer.NextToken()

		if nextToken.Type != tokenType.expectedType {
			assert.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", index, tokenType.expectedType, nextToken.Type)
		}

		if nextToken.Literal != tokenType.expectedLiteral {
			assert.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", index, tokenType.expectedLiteral, nextToken.Literal)
		}
	}
}
