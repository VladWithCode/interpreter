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
	checkParserErrors(t, p)
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

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New([]byte(input))
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d\n", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got %T", stmt)
			continue
		}
		if !bytes.Equal(returnStmt.TokenLiteral(), []byte("return")) {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name []byte) bool {
	if !bytes.Equal(stmt.TokenLiteral(), []byte("let")) {
		t.Errorf("s.TokenLiteral not \"let\". got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if !bytes.Equal(letStmt.Name.Value, name) {
		t.Errorf("letStmt.Name.Value not '%s'. got %q", name, letStmt.Name.Value)
		return false
	}
	if !bytes.Equal(letStmt.Name.TokenLiteral(), name) {
		t.Errorf("stmt.Name not '%s'. got %q", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
