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
	StringLiteralNode  NodeType = "StringLiteral"
	AssignmentNode     NodeType = "Assignment"
	BooleanLiteralNode NodeType = "BooleanLiteral"
	IfStmtNode         NodeType = "IfStmt"
	BlockStmtNode      NodeType = "BlockStmt"
	WhileStmtNode      NodeType = "WhileStmt"
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

type StringLiteral struct {
	Type  NodeType
	Value string
}

func (s *StringLiteral) GetType() NodeType { return s.Type }
func (s *StringLiteral) expressionNode()   {}

type BooleanLiteral struct {
	Type  NodeType
	Value bool
}

func (b *BooleanLiteral) GetType() NodeType { return b.Type }
func (b *BooleanLiteral) expressionNode()   {}

type VarDeclaration struct {
	Type       NodeType
	Identifier string
	Value      Expr
}

func (v *VarDeclaration) GetType() NodeType { return v.Type }

type Assignment struct {
	Type       NodeType
	Identifier string
	Value      Expr
}

func (a *Assignment) GetType() NodeType { return a.Type }

type BlockStmt struct {
	Type NodeType
	Body []Stmt
}

func (b *BlockStmt) GetType() NodeType { return b.Type }

type IfStmt struct {
	Type       NodeType
	Condition  Expr
	ThenBranch *BlockStmt
	ElseBranch *BlockStmt
}

func (i *IfStmt) GetType() NodeType { return i.Type }

type WhileStmt struct {
	Type      NodeType
	Condition Expr
	Body      *BlockStmt
}

func (w *WhileStmt) GetType() NodeType { return w.Type }

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
	case *StringLiteral:
		fmt.Printf("%sStringLiteral: %s\n", indent, n.Value)
	case *Assignment:
		fmt.Printf("%sAssignment: %s\n", indent, n.Identifier)
		printAST(n.Value, indent+"  ")
	case *BooleanLiteral:
		fmt.Printf("%sBooleanLiteral: %t\n", indent, n.Value)
	case *BlockStmt:
		fmt.Printf("%sBlockStmt\n", indent)
		for _, stmt := range n.Body {
			printAST(stmt, indent+"  ")
		}
	case *IfStmt:
		fmt.Printf("%sIfStmt\n", indent)
		fmt.Printf("%sCondition:\n", indent+"  ")
		printAST(n.Condition, indent+"    ")
		fmt.Printf("%sThenBranch:\n", indent+"  ")
		printAST(n.ThenBranch, indent+"    ")
		if n.ElseBranch != nil {
			fmt.Printf("%sElseBranch:\n", indent+"  ")
			printAST(n.ElseBranch, indent+"    ")
		}
	case *WhileStmt:
		fmt.Println(indent + "WhileStmt")
		fmt.Println(indent + "  Condition:")
		printAST(n.Condition, indent+"    ")
		fmt.Println(indent + "  Body:")
		printAST(n.Body, indent+"    ")
	default:
		fmt.Printf("%sUnknown node type: %T\n", indent, n)
	}
}
