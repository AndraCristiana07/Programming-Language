package testing

import (
	"my_language/ast"
	"my_language/parser"

	"github.com/antlr4-go/antlr/v4"
)

func runInterpreter(inputCode string) *ast.Environment {
	inputStream := antlr.NewInputStream(inputCode)
	lexer := parser.NewGrammarLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGrammarParser(tokenStream)

	tree := p.Program()

	visitor := ast.NewVisitor()

	tree.Accept(visitor)

	return visitor.GetEnvironment()
}
