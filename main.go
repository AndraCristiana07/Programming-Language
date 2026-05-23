package main

import (
	"fmt"
	"my_language/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputCode := `
		var x = 10 * 5
		x = x + 1`

	input := antlr.NewInputStream(inputCode)

	lexer := parser.NewGrammarLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGrammarParser(tokens)

	tree := p.Program()

	fmt.Println("--- Parse Tree ---")
	fmt.Println(tree.ToStringTree(nil, p))

	eval := NewVisitor()
	tree.Accept(eval)

	fmt.Println("Final variable states:")
	for k, v := range eval.vars {
		fmt.Printf("%s = %v\n", k, v)
	}
}
