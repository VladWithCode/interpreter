package token

import "fmt"

type TokenType int

const (
	EOF TokenType = iota
	ILLEGAL

	// Identifiers + literals
	IDENT
	INT

	// Operators
	ASSIGN
	PLUS
	MINUS
	BANG
	SLASH
	ASTERISK
	LT
	GT
	EQ
	NOT_EQ

	// Delimiters
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	ELSEIF
	RETURN
)

func (tt TokenType) String() string {
	if typeStr, ok := TokenTypeMap[tt]; ok {
		return typeStr
	}

	return fmt.Sprintf("%d", tt)
}

var TokenTypeMap = map[TokenType]string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	// Identifiers + literals
	IDENT: "IDENT",
	INT:   "INT",
	// Operators
	ASSIGN:   "ASSIGN",
	PLUS:     "PLUS",
	MINUS:    "MINUS",
	BANG:     "BANG",
	ASTERISK: "ASTERISK",
	SLASH:    "SLASH",
	LT:       "LT",
	GT:       "GT",
	EQ:       "EQ",
	NOT_EQ:   "NOT_EQ",
	// Delimiters
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
	LPAREN:    "LPAREN",
	RPAREN:    "RPAREN",
	LBRACE:    "LBRACE",
	RBRACE:    "RBRACE",
	// Keywords
	FUNCTION: "FUNCTION",
	LET:      "LET",
	TRUE:     "TRUE",
	FALSE:    "FLASE",
	IF:       "IF",
	ELSE:     "ELSE",
	ELSEIF:   "ELSEIF",
	RETURN:   "RETURN",
}

type Token struct {
	Type    TokenType
	Literal []byte
}

func NewToken(tokenType TokenType, ch []byte) *Token {
	return &Token{
		Type:    tokenType,
		Literal: ch,
	}
}

func TokenTypeToString(tokenType TokenType) string {
	if typeStr, ok := TokenTypeMap[tokenType]; ok {
		return typeStr
	}

	return fmt.Sprintf("%d", tokenType)
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"elseif": ELSEIF,
	"return": RETURN,
}

func LookupIdent(ident []byte) TokenType {
	strIdent := string(ident)
	if tok, ok := keywords[strIdent]; ok {
		return tok
	}
	return IDENT
}
