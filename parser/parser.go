package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lex *lexer.Lexer // current line which we are evaluating

	currentToken token.Token
	peekToken    token.Token // next token
	error        []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lex: l, error: []string{}}

	// Read two tokens so current and next tokens are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.error
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)

	p.error = append(p.error, msg)
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) isCurrentToken(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) expectedPeek(t token.TokenType) bool {
	if p.isPeekToken(t) {
		p.nextToken()
		return true
	} else {
		// if the next token is not the expected token
		// then we will add an error message to the error list
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}

	// token.IDENT is the name of the variable
	// let Variable_Name = expression
	// after let the next should be a variable name
	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	// checking if the next token is an assignment token or not
	// let Variable_Name '=' expression
	if !p.expectedPeek(token.ASSSIGN) {
		return nil
	}

	// keep on iterating until we reach the end of the statement
	// which is a semicolon
	// ex: let x = 5;
	for !p.isCurrentToken(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	// will keep on looping until we reach the |;| symbol
	for !p.isCurrentToken(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	// creating the root of the ast tree
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		// advances both p.curToken and p.peekToken
		p.nextToken()
	}

	return program
}
