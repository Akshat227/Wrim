package lexer

import (
	"Wrim/token"
	"fmt"
)

type Lexer struct {
	// input is the full source being lexed.
	input string

	// position is the index of the current character in input.
	position int

	// readPosition is the index where we will read next — it is
	// always position+1 except at EOF. This allows simple 1-char
	// lookahead without slicing the input repeatedly.
	readPosition int

	// ch holds the current character under examination. When ch == 0
	// it means we've reached EOF.
	ch byte

	// Debug enables printing of each token produced. Useful for
	// development and maintainability; callers can enable it via
	// `l.SetDebug(true)` after creating the lexer.
	Debug bool
}

// Initialize a lexer with input source
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Advance one character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// Produce the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			// Read an identifier or keyword starting at the current
			// position. `readIdentifier` will advance `l` until a
			// non-letter is found.
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			// Print token if debugging is enabled.
			if l.Debug {
				fmt.Printf("LEXER -> Token: Type=%s Literal=%q\n", tok.Type, tok.Literal)
			}
			return tok
		} else if isDigit(l.ch) {
			// Read a contiguous sequence of digits as an integer
			// literal. `readNumber` advances the lexer's position.
			tok.Type = token.INT
			tok.Literal = l.readNumber()

			// Print token if debugging is enabled.
			if l.Debug {
				fmt.Printf("LEXER -> Token: Type=%s Literal=%q\n", tok.Type, tok.Literal)
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Print token if debugging is enabled before we advance to the
	// next character.
	if l.Debug {
		fmt.Printf("LEXER -> Token: Type=%s Literal=%q\n", tok.Type, tok.Literal)
	}

	l.readChar()
	return tok
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{Type: t, Literal: string(ch)}
}

// SetDebug toggles token printing for this lexer instance.
func (l *Lexer) SetDebug(d bool) {
	l.Debug = d
}

// Helpers ----------------------------------------------------------

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
