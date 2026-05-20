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

		var a = 10
		print "a is now: " + a
		a = a + 5
		print "a is now: " + a
		a = "is string now"
		print "a is now: " + a

		var target = 20
        print target > 15
        print target == 10 + 10
        print 5 != 5
        var truth = true
        print truth
	
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
