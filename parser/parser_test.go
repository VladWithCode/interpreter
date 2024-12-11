package parser

import (
	"bytes"
	"testing"

	"github.com/vladwithcode/monkey-interpreter/ast"
	"github.com/vladwithcode/monkey-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := []byte(`
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`)

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() produced nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d\n", len(program.Statements))
	}

	test := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range test {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, []byte(tt.expectedIdentifier)) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name []byte) bool {
	if bytes.Equal(stmt.TokenLiteral(), []byte("let")) {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if bytes.Equal(letStmt.Name.Value, name) {
		t.Errorf("letStmt.Name.Value not '%s'. got%s", name, letStmt.Name.Value)
		return false
	}
	if bytes.Equal(letStmt.Name.TokenLiteral(), name) {
		t.Errorf("stmt.Name not '%s'. got%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
