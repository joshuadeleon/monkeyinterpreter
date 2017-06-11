package lexer

import "monkeyinterpreter/token"

// Lexer structure encapulates a lexer
type Lexer struct {
	input        string
	position     int
	readPosition int
	value        byte
}

// New instantiates new lexer
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readCharacter()
	return lexer
}

// NextToken advances the read position and returns the next token
func (lexer *Lexer) NextToken() token.Token {
	var nextToken token.Token

	switch lexer.value {
	case '=':
		nextToken = newToken(token.ASSIGN, lexer.value)
	case ';':
		nextToken = newToken(token.SEMICOLON, lexer.value)
	case '(':
		nextToken = newToken(token.LPAREN, lexer.value)
	case ')':
		nextToken = newToken(token.RPAREN, lexer.value)
	case '{':
		nextToken = newToken(token.LBRACE, lexer.value)
	case '}':
		nextToken = newToken(token.RBRACE, lexer.value)
	case ',':
		nextToken = newToken(token.COMMA, lexer.value)
	case '+':
		nextToken = newToken(token.PLUS, lexer.value)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	}

	lexer.readCharacter()
	return nextToken
}

// newToken instantiates a Token from the parameters
func newToken(tokenType token.Type, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

// readCharacter reads the next character and advances position of the pointers
func (lexer *Lexer) readCharacter() {
	lexer.value = lexer.getNextCharacter()
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

// getNextCharacter returns the next character
func (lexer *Lexer) getNextCharacter() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}

	return lexer.input[lexer.readPosition]
}
