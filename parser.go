package main

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

// read toekns and build AST
func (p *Parser) Parse() *Program {
	program := &Program{
		Type: ProgramNode,
		Body: []Stmt{},
	}

	// do it till end of toekns
	for !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Body = append(program.Body, stmt)
		}
	}
	return program
}

// is at end
func (p *Parser) isAtEnd() bool {
	return p.pos >= len(p.tokens)
}

// peek current token
func (p *Parser) peek() Token {
	if p.isAtEnd() {
		return Token{Type: "EOF", Value: ""}
	}
	return p.tokens[p.pos]
}

// advance to next token
func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.pos++
	}
	return p.tokens[p.pos-1]
}

// check if current token matches expected type
func (p *Parser) match(expectedType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	if p.peek().Type == expectedType {
		p.advance()
		return true
	}
	return false
}

// expect current token to be of expected type, otherwise panic
func (p *Parser) expect(expectedType TokenType) Token {
	if p.isAtEnd() {
		panic("Unexpected end of input")
	}
	if p.peek().Type == expectedType {
		return p.advance()
	}
	panic("Unexpected token: " + p.peek().Value)
}

// parse statement
func (p *Parser) parseStatement() Stmt {
	if p.match(VarToken) {
		return p.parseVarDeclaration()
	}
	if p.match(PrintToken) {
		return p.parsePrintStatement()
	}
	panic("Unexpected token: " + p.peek().Value)
}

// parse var
func (p *Parser) parseVarDeclaration() *VarDeclaration {
	// p.expect(VarToken) //  "var"
	identifierToken := p.expect(IdentifierToken)
	p.expect(EqualsToken) //  "="
	value := p.parseExpression()
	return &VarDeclaration{
		Type:       VarDeclarationNode,
		Identifier: identifierToken.Value,
		Value:      value,
	}
}

// parse print
func (p *Parser) parsePrintStatement() *PrintStmt {
	value := p.parseExpression()
	return &PrintStmt{
		Type:  PrintStmtNode,
		Value: value,
	}
}

// parse expression
func (p *Parser) parseExpression() Expr {
	return p.parseAddition()
}

func (p *Parser) parseAddition() Expr {
	// call highest lvl first
	left := p.parseMultiplication()

	for p.match(PlusToken) || p.match(MinusToken) {
		operator := p.tokens[p.pos-1].Value
		right := p.parseMultiplication() // call next lvl for right side

		left = &BinaryExpr{
			Type:     BinaryExprNode,
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}

	return left
}

func (p *Parser) parseMultiplication() Expr {
	left := p.parsePrimary()
	for p.match(StarToken) || p.match(SlashToken) {
		operator := p.tokens[p.pos-1].Value
		right := p.parsePrimary()
		left = &BinaryExpr{
			Type:     BinaryExprNode,
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}
	return left
}

// parse leaf nodes (identifiers, numbers)
func (p *Parser) parsePrimary() Expr {
	if p.match(IdentifierToken) {
		return &Identifier{
			Type:   IdentifierNode,
			Symbol: p.tokens[p.pos-1].Value,
		}
	} else if p.match(NumberToken) {
		return &NumericLiteral{
			Type:  NumericLiteralNode,
			Value: atoi(p.tokens[p.pos-1].Value),
		}
	} else if p.match(StringToken) {
		return &StringLiteral{
			Type:  StringLiteralNode,
			Value: p.tokens[p.pos-1].Value,
		}
	} else {
		panic("Unexpected token: " + p.peek().Value)
	}
}

// helper to convert string to int
func atoi(s string) int {
	result := 0
	for _, char := range s {
		result = result*10 + int(char-'0')
	}
	return result
}
