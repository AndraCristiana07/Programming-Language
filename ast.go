package main

import "fmt"

type NodeType string

const (
	ProgramNode        NodeType = "Program"
	VarDeclarationNode NodeType = "VarDeclaration"
	NumericLiteralNode NodeType = "NumericLiteral"
	IdentifierNode     NodeType = "Identifier"
	BinaryExprNode     NodeType = "BinaryExpr"
	PrintStmtNode      NodeType = "PrintStmt"
)

type Stmt interface {
	GetType() NodeType
}

type Expr interface {
	Stmt
	expressionNode()
}

type Program struct {
	Type NodeType
	Body []Stmt
}

type PrintStmt struct {
	Type  NodeType
	Value Expr
}

func (p *PrintStmt) GetType() NodeType { return p.Type }
func (p *Program) GetType() NodeType   { return p.Type }

type BinaryExpr struct {
	Type     NodeType
	Left     Expr
	Right    Expr
	Operator string // "+", "-", "*", "/"
}

func (b *BinaryExpr) GetType() NodeType { return b.Type }
func (b *BinaryExpr) expressionNode()   {}

type Identifier struct {
	Type   NodeType
	Symbol string
}

func (i *Identifier) GetType() NodeType { return i.Type }
func (i *Identifier) expressionNode()   {}

type NumericLiteral struct {
	Type  NodeType
	Value int
}

func (n *NumericLiteral) GetType() NodeType { return n.Type }
func (n *NumericLiteral) expressionNode()   {}

type VarDeclaration struct {
	Type       NodeType
	Identifier string
	Value      Expr
}

func (v *VarDeclaration) GetType() NodeType { return v.Type }

func printAST(node Stmt, indent string) {
	switch n := node.(type) {
	case *Program:
		fmt.Println(indent + "Program")
		for _, stmt := range n.Body {
			printAST(stmt, indent+"  ")
		}
	case *VarDeclaration:
		fmt.Printf("%sVarDeclaration: %s\n", indent, n.Identifier)
		printAST(n.Value, indent+"  ")
	case *PrintStmt:
		fmt.Printf("%sPrintStmt\n", indent)
		printAST(n.Value, indent+"  ")
	case *BinaryExpr:
		fmt.Printf("%sBinaryExpr: %s\n", indent, n.Operator)
		printAST(n.Left, indent+"  ")
		printAST(n.Right, indent+"  ")
	case *Identifier:
		fmt.Printf("%sIdentifier: %s\n", indent, n.Symbol)
	case *NumericLiteral:
		fmt.Printf("%sNumericLiteral: %d\n", indent, n.Value)
	default:
		fmt.Printf("%sUnknown node type: %T\n", indent, n)
	}
}
