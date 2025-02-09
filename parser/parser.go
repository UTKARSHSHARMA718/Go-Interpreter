package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lex *lexer.Lexer // current line which we are evaluating

	currentToken token.Token
	peekToken    token.Token // next token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lex: l}

	// Read two tokens so current and next tokens are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken }

	// token.IDENT is the name of the variable
	if !p.expectedPeek(token.IDENT){
		return nil
	}

	statement.Name = &ast.Identifier{Token: }
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parserLetStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	//creating the root of the ast tree
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
