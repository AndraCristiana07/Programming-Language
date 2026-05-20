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

		var b = 10 <= 5
		print b
		var c = 10 >= 5
		print c
		print 10 >= 9

		if 10 > 3 {
			print "bigger"
		} else {
			print "smaller"
		}

		var f = 10
		print "f:" + f
		if true {
			var f = 50
			print "f:" + f
			
		}
		print "f:" + f

		var d = 10
		print d
		if true {
			d = 50
			print d
			var e = 20
			print e
		}
		print d
		
		var g = 1
		while g <= 5 {
			print "loop step: " + g
			g = g + 1
		}


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
