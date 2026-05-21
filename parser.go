package main

import "fmt"

// TODO: make boolean assertion possible b = true; if b == true: ...
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

func (p *Parser) lookAhead(offset int) Token {
	if p.pos+offset >= len(p.tokens) {
		return Token{Type: "EOF", Value: ""}
	}
	return p.tokens[p.pos+offset]
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
	currentToken := p.peek()
	if currentToken.Type == expectedType {
		return p.advance()
	}
	panic(fmt.Sprintf("Syntax Error on Line %d: Expected token type '%s' but found '%s' with value '%s'",
		currentToken.Line, expectedType, currentToken.Type, currentToken.Value))
}

// parse statement
func (p *Parser) parseStatement() Stmt {
	if p.match(VarToken) {
		return p.parseVarDeclaration()
	}
	if p.match(PrintToken) {
		return p.parsePrintStatement()
	}
	if p.match(IfToken) {
		return p.parseIfStatement()
	}
	if p.match(WhileToken) {
		return p.parseWhileStatement()
	}
	if p.match(ForToken) {
		return p.parseForStatement()
	}
	if p.match(FuncToken) {
		return p.parseFunctionDeclaration()
	}
	if p.match(ReturnToken) {
		return p.parseReturnStatement()
	}
	// check for function invocations
	if p.peek().Type == IdentifierToken && p.lookAhead(1).Type == LParenToken {
		return p.parseExpression()
	}

	if p.peek().Type == IdentifierToken && p.lookAhead(1).Type == EqualsToken {
		return p.parseAssignment()
	}
	panic(fmt.Sprintf("Syntax Error on Line %d: Unexpected token: %s", p.peek().Line, p.peek().Value))
}

// for var i = 1; i <= 3; i = i + 1 {
//     print i
// }

// var i = 1
//
//	while i <= 3 {
//	    print i
//	    i = i + 1
//	}
func (p *Parser) parseForStatement() Stmt {
	// parse the initializer (var i = 1)
	var initializer Stmt = nil
	if p.match(VarToken) {
		initializer = p.parseVarDeclaration()
	} else {
		// if not var declaration, check if direct assignment like i = 1
		if p.peek().Type == IdentifierToken && p.lookAhead(1).Type == EqualsToken {
			initializer = p.parseAssignment()
		}
	}
	p.expect(SemiColonToken)

	// parse the condition  (i <= 3)
	condition := p.parseExpression()
	p.expect(SemiColonToken)

	// parse the step (i = i + 1)
	increment := p.parseAssignment()

	// parse the loop body ({print i })
	p.expect(LBraceToken)
	loopBody := p.parseBlockStatement()

	// inc clause at the end of the loop's statement
	loopBody.Body = append(loopBody.Body, increment)

	whileLoop := &WhileStmt{
		Type:      WhileStmtNode,
		Condition: condition,
		Body:      loopBody,
	}

	// make parent block to hold initializer and while loop
	parentBlock := &BlockStmt{
		Type: BlockStmtNode,
		Body: []Stmt{initializer, whileLoop},
	}

	return parentBlock
}

func (p *Parser) parseFunctionDeclaration() *FuncDecl {
	nameToken := p.expect(IdentifierToken)

	p.expect(LParenToken) // expect '('
	var params []string
	if p.peek().Type != RParenToken {
		for {
			paramToken := p.expect(IdentifierToken)
			params = append(params, paramToken.Value)
			if !p.match(CommaToken) {
				break
			}
		}
	}
	p.expect(RParenToken) // expect ')'

	p.expect(LBraceToken) // expect '{'
	body := p.parseBlockStatement()

	return &FuncDecl{
		Type:       FuncDeclNode,
		Name:       nameToken.Value,
		Parameters: params,
		Body:       body,
	}
}

func (p *Parser) parseReturnStatement() *ReturnStmt {
	var value Expr = nil
	// if next token is not closing brce or end of block -> parse the return expr
	if p.peek().Type != RBraceToken {
		value = p.parseExpression()
	}
	return &ReturnStmt{
		Type:  RreturnStmtNode,
		Value: value,
	}
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

// parse assignment
func (p *Parser) parseAssignment() *Assignment {
	// consume identifier and equals
	identifierToken := p.expect(IdentifierToken)
	p.expect(EqualsToken)
	value := p.parseExpression()
	return &Assignment{
		Type:       AssignmentNode,
		Identifier: identifierToken.Value,
		Value:      value,
	}
}

// parse if statement
func (p *Parser) parseIfStatement() *IfStmt {
	condition := p.parseExpression()

	p.expect(LBraceToken) // consume "{"
	thenBranch := p.parseBlockStatement()

	var elseBranch *BlockStmt = nil

	// handle optional else branch
	if p.match(ElseToken) {
		p.expect(LBraceToken)
		elseBranch = p.parseBlockStatement()
	}

	return &IfStmt{
		Type:       IfStmtNode,
		Condition:  condition,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
	}
}

func (p *Parser) parseWhileStatement() *WhileStmt {
	condition := p.parseExpression()

	p.expect(LBraceToken)
	body := p.parseBlockStatement()

	return &WhileStmt{
		Type:      WhileStmtNode,
		Condition: condition,
		Body:      body,
	}
}

// parse block statement
func (p *Parser) parseBlockStatement() *BlockStmt {
	var body []Stmt

	// if block starts with "{" we consume it
	for !p.isAtEnd() && p.peek().Type != RBraceToken {
		body = append(body, p.parseStatement())
	}

	p.expect(RBraceToken) // consume '}'

	return &BlockStmt{
		Type: BlockStmtNode,
		Body: body,
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
	return p.parseComparison()
}

func (p *Parser) parseComparison() Expr {
	left := p.parseAddition() // math first

	for p.match(LessToken) || p.match(GreaterToken) || p.match(EqualEqualToken) || p.match(BangEqualToken) || p.match(LessEqualToken) || p.match(GreaterEqualToken) {
		operator := p.tokens[p.pos-1].Value
		right := p.parseAddition()

		left = &BinaryExpr{
			Type:     BinaryExprNode,
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}

	return left
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
		name := p.tokens[p.pos-1].Value

		// check if the identifier is followed by a '('
		if p.match(LParenToken) {
			var args []Expr
			if p.peek().Type != RParenToken {
				for {
					args = append(args, p.parseExpression())
					if !p.match(CommaToken) {
						break
					}
				}
			}
			p.expect(RParenToken)

			return &CallExpr{
				Type:      CallExprNode,
				Callee:    name,
				Arguments: args,
			}
		}

		return &Identifier{Type: IdentifierNode, Symbol: name}
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
	} else if p.match(TrueToken) {
		return &BooleanLiteral{
			Type:  BooleanLiteralNode,
			Value: true,
		}
	} else if p.match(FalseToken) {
		return &BooleanLiteral{
			Type:  BooleanLiteralNode,
			Value: false,
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
