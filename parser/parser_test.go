package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetParser(t *testing.T) {
	input := `let a = 1;
	let b= 12;
	let feynman=26341896;`

	l := lexer.New(input) // to read the input code line by line
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

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
		{"b"},
		{"feynman"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		// wont execute futher once found any error
		if !testLetStatements(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnParser(t *testing.T) {
	input := `
return 5;
 return 10;
 return 993322;`

	l := lexer.New(input) // to read the input code line by line
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParserProgram returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, statement := range program.Statements {

		fmt.Println("---------------", statement)
		returnStatement, ok := statement.(*ast.ReturnStatement)
		fmt.Println("---------------", returnStatement, ok)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", statement)
			continue
		}

		if returnStatement.ReturnValue == nil {
			t.Errorf("returnStmt.ReturnValue is nil")
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			fmt.Println("FINISHED")
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStatement.TokenLiteral())
		}
	}

}

func testLetStatements(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	// checking does the statement contains a specific type 'LetStatement'
	letsmt, ok := statement.(*ast.LetStatement)

	if !ok {
		t.Errorf("statement not *ast.LetStatement. got=%T", statement)
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
