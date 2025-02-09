package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetParser(t *testing.T) {
	input := `let a = 1;
	
	let b = 12;
	let feynman = 26341896;`

	l := lexer.New(input) // to read the input code line by line
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParserProgram returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"y"},
		{"feynman"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	// checking does the statement contains a specific type 'LetStatement'
	letsmt, ok := statement.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", statement)
		return false
	}

	if letsmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letsmt.Name.Value)
		return false
	}

	if letsmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letsmt.Name)
		return false
	}

	return true
}
