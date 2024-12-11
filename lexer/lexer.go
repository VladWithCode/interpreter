package lexer

import (
	"github.com/vladwithcode/monkey-interpreter/token"
)

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) HasNext() bool {
	return l.position < len(l.input)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = '0'
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return '0'
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() *token.Token {
	var tok *token.Token
	l.skipWhiteSpace()

	ch := []byte{l.ch}
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch = append(ch, l.input[l.readPosition])
			tok = token.NewToken(token.EQ, ch)
			l.readChar()
		} else {
			tok = token.NewToken(token.ASSIGN, ch)
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, ch)
	case '(':
		tok = token.NewToken(token.LPAREN, ch)
	case ')':
		tok = token.NewToken(token.RPAREN, ch)
	case ',':
		tok = token.NewToken(token.COMMA, ch)
	case '+':
		tok = token.NewToken(token.PLUS, ch)
	case '-':
		tok = token.NewToken(token.MINUS, ch)
	case '!':
		if l.peekChar() == '=' {
			ch = append(ch, l.input[l.readPosition])
			tok = token.NewToken(token.NOT_EQ, ch)
			l.readChar()
		} else {
			tok = token.NewToken(token.BANG, ch)
		}
	case '*':
		tok = token.NewToken(token.ASTERISK, ch)
	case '/':
		tok = token.NewToken(token.SLASH, ch)
	case '<':
		tok = token.NewToken(token.LT, ch)
	case '>':
		tok = token.NewToken(token.GT, ch)
	case '{':
		tok = token.NewToken(token.LBRACE, ch)
	case '}':
		tok = token.NewToken(token.RBRACE, ch)
	case '0':
		tok = token.NewToken(token.EOF, ch)
	default:
		tok = &token.Token{}
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
		} else if isDigit(l.ch) {
			tok.Literal = l.readInt()
			tok.Type = token.INT
		} else {
			tok = token.NewToken(token.ILLEGAL, ch)
		}

		return tok
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() []byte {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readInt() []byte {
	pos := l.position
	for isDigit(l.ch) && l.HasNext() {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func New(input []byte) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
