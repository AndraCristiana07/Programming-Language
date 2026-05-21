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
	FuncDeclNode       NodeType = "FuncDecl"
	CallExprNode       NodeType = "CallExpr"
	RreturnStmtNode    NodeType = "ReturnStmt"
	UnaryExprNode      NodeType = "UnaryExpr"
	ArrayLiteralNode   NodeType = "ArrayLiteral"
	IndexExprNode      NodeType = "IndexExpr"
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

type UnaryExpr struct {
	Type     NodeType
	Operator string // "not"
	Right    Expr
}

func (u *UnaryExpr) GetType() NodeType { return u.Type }
func (u *UnaryExpr) expressionNode()   {}

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

type FuncDecl struct {
	Type       NodeType
	Name       string
	Parameters []string
	Body       *BlockStmt
}

func (f *FuncDecl) GetType() NodeType { return f.Type }

type CallExpr struct {
	Type      NodeType
	Callee    string // function name being called
	Arguments []Expr
}

func (c *CallExpr) GetType() NodeType { return c.Type }
func (c *CallExpr) expressionNode()   {}

type ReturnStmt struct {
	Type  NodeType
	Value Expr
}

func (r *ReturnStmt) GetType() NodeType { return r.Type }

// a literal array initialization like [1, 2, 3]
type ArrayLiteral struct {
	Type     NodeType
	Elements []Expr
}

// for sreading an index item like items[0]
type IndexExpr struct {
	Type  NodeType
	Left  Expr // array object
	Index Expr // index value expression inside the brackets
}

func (a *ArrayLiteral) GetType() NodeType { return a.Type }
func (a *ArrayLiteral) expressionNode()   {}
func (i *IndexExpr) GetType() NodeType    { return i.Type }
func (i *IndexExpr) expressionNode()      {}

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

	case *FuncDecl:
		fmt.Printf("%sFuncDecl: %s(%v)\n", indent, n.Name, n.Parameters)
		printAST(n.Body, indent+"  ")

	case *CallExpr:
		fmt.Printf("%sCallExpr: %s\n", indent, n.Callee)
		for _, arg := range n.Arguments {
			printAST(arg, indent+"  ")
		}
	case *ReturnStmt:
		fmt.Printf("%sReturnStmt\n", indent)
		printAST(n.Value, indent+"  ")

	default:
		fmt.Printf("%sUnknown node type: %T\n", indent, n)
	}
}
