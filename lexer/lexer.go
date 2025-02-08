package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	in := &Lexer{input: input}
	in.readChar()
	return in
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // if string input is finished
		return
	}

	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenValue token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenValue, Literal: string(ch)} // creating and return a struct of type Token
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// attaching a function to Lexer struct
func (l *Lexer) readIdentifier() string {
	position := l.position // current position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // equivalent to str.substring(position, l.position) method
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // equivalent to str.substring(position, l.position) method
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token // declare a variable of type Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok.Literal = "=="
			tok.Type = token.DOUBLE_EQ
			l.readChar() // move to next character
			l.readChar() // move to next character
			return tok
		}
		tok = newToken(token.ASSSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
			l.readChar()
			l.readChar()
			return tok
		}
		tok = newToken(token.BANG, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case 0: // null value which we get at end of file
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() //readIdentifier() does exactly what its name suggests: it reads in an identifier and advances our lexer’s positions until it encounters a non-letter-character.
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
