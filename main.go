package main

import "fmt"

func main() {
	input := `
		var x = 10 + 2
		var y = 10 * 2 + 5
		var z = x + y
		print x
		print y
		print z

		var greeting = "Hello, "
		var name = "World!"
		print greeting + name
		print "X is " + x
	
	`
	tokens := Lex(input)

	for _, t := range tokens {
		fmt.Printf("{Type: %-12s Value: %q}\n", t.Type, t.Value)
	}

	parser := NewParser(tokens)
	ast := parser.Parse()
	printAST(ast, "")

	env := NewEnvironment()
	Eval(ast, env)

}
