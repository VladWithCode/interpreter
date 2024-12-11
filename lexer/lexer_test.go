package lexer_test

import (
	"bytes"
	"testing"

	. "github.com/vladwithcode/monkey-interpreter/lexer"
	"github.com/vladwithcode/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := []byte(`let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		10 != 9;
		10 != 10

		`)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral []byte
	}{
		{token.LET, []byte("let")},
		{token.IDENT, []byte("five")},
		{token.ASSIGN, []byte{'='}},
		{token.INT, []byte{'5'}},
		{token.SEMICOLON, []byte{';'}},
		{token.LET, []byte("let")},
		{token.IDENT, []byte("ten")},
		{token.ASSIGN, []byte{'='}},
		{token.INT, []byte("10")},
		{token.SEMICOLON, []byte{';'}},
		{token.LET, []byte("let")},
		{token.IDENT, []byte("add")},
		{token.ASSIGN, []byte{'='}},
		{token.FUNCTION, []byte("fn")},
		{token.LPAREN, []byte{'('}},
		{token.IDENT, []byte{'x'}},
		{token.COMMA, []byte{','}},
		{token.IDENT, []byte{'y'}},
		{token.RPAREN, []byte{')'}},
		{token.LBRACE, []byte{'{'}},
		{token.IDENT, []byte{'x'}},
		{token.PLUS, []byte{'+'}},
		{token.IDENT, []byte{'y'}},
		{token.SEMICOLON, []byte{';'}},
		{token.RBRACE, []byte{'}'}},
		{token.SEMICOLON, []byte{';'}},
		{token.LET, []byte("let")},
		{token.IDENT, []byte("result")},
		{token.ASSIGN, []byte{'='}},
		{token.IDENT, []byte("add")},
		{token.LPAREN, []byte{'('}},
		{token.IDENT, []byte("five")},
		{token.COMMA, []byte{','}},
		{token.IDENT, []byte("ten")},
		{token.RPAREN, []byte{')'}},
		{token.SEMICOLON, []byte{';'}},
		{token.BANG, []byte{'!'}},
		{token.MINUS, []byte{'-'}},
		{token.SLASH, []byte{'/'}},
		{token.ASTERISK, []byte{'*'}},
		{token.INT, []byte{'5'}},
		{token.SEMICOLON, []byte{';'}},
		{token.INT, []byte{'5'}},
		{token.LT, []byte{'<'}},
		{token.INT, []byte("10")},
		{token.GT, []byte{'>'}},
		{token.INT, []byte{'5'}},
		{token.SEMICOLON, []byte{';'}},
		{token.IF, []byte("if")},
		{token.LPAREN, []byte{'('}},
		{token.INT, []byte{'5'}},
		{token.LT, []byte{'<'}},
		{token.INT, []byte("10")},
		{token.RPAREN, []byte{')'}},
		{token.LBRACE, []byte{'{'}},
		{token.RETURN, []byte("return")},
		{token.TRUE, []byte("true")},
		{token.SEMICOLON, []byte{';'}},
		{token.RBRACE, []byte{'}'}},
		{token.ELSE, []byte("else")},
		{token.LBRACE, []byte{'{'}},
		{token.RETURN, []byte("return")},
		{token.FALSE, []byte("false")},
		{token.SEMICOLON, []byte{';'}},
		{token.RBRACE, []byte{'}'}},
		{token.INT, []byte("10")},
		{token.EQ, []byte("==")},
		{token.INT, []byte("10")},
		{token.SEMICOLON, []byte{';'}},
		{token.INT, []byte("10")},
		{token.NOT_EQ, []byte("!=")},
		{token.INT, []byte{'9'}},
		{token.SEMICOLON, []byte{';'}},
		{token.INT, []byte("10")},
		{token.NOT_EQ, []byte("!=")},
		{token.INT, []byte("10")},
		{token.EOF, []byte{'0'}},
	}

	l := New(input)

	for i, tt := range tests {
		actualToken := l.NextToken()
		if actualToken.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				token.TokenTypeToString(tt.expectedType),
				token.TokenTypeToString(actualToken.Type),
			)
		}

		if !bytes.Equal(actualToken.Literal, tt.expectedLiteral) {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i,
				tt.expectedLiteral,
				actualToken.Literal,
			)
		}
	}
}
