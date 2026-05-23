package main

import (
	"fmt"
	"my_language/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputCode := `
		var x = 10 * 5
		print x
		x = x + 1

		print "Hello"

		var score = 50 + 1
    	print "final score: " + score

		var isGreater = 10 > 5
		var checkEquality = isGreater == true
		var uneqTest = 5 != 3
		var lesseqTest = 5 <= 5
		var greateqTest = 5 >= 4

		var isPassed = false
		if (score >= 50) {
			isPassed = true
		} else {
			isPassed = false
		}
		print "Passed: " + isPassed
		
		var i = 0
		while ( i < 5 ) {
			print i
			i = i + 1
		}
		print "i is: " + i
		`

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
