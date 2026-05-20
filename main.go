package main

import "fmt"

func main() {
	input := "var x = 10 + 20"
	tokens := Lex(input)

	for _, t := range tokens {
		fmt.Printf("{Type: %-12s Value: %q}\n", t.Type, t.Value)
	}

	parser := NewParser(tokens)
	ast := parser.Parse()
	printAST(ast, "")

	env := NewEnvironment()
	result := Eval(ast, env)
	fmt.Printf("x = %d\n", result)
}
