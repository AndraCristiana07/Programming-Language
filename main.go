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

		print "Starting for loop:"
        for var i = 1; i <= 3; i = i + 1 {
            print "i is: " + i
        }
        print "Loop complete"

		print "Starting for loop2:"
        for var i = 6; i >= 2; i = i - 1 {
            print "i is: " + i
        }
        print "Loop2 complete"

		var i = 2
		print "Starting for loop3:"
		for i = 3; i < 6; i = i + 1 {
			print "i is:" + i
		}
		print "Loop3 complete"

		func myfn(user) {
			var i = 1
			while i < 4 {
				print "Hi " + user + " at i: " + i
				i = i + 1
			}
		}

		var name = "Ana"
		myfn(name)

		func add(a, b) {
			return a + b
		}
		var sum = add(10, 20)
		print "sum is: " + sum

		var count = 0
		func printsmth(a) {
			var i = 5
			while i <= 10 {
				a = a + 2
				i = i + 1
			}
			return a
		}
		print printsmth(count)

		print 10 % 3
		print 2 ** 3


		var counter = 10
		counter--
		print "Counter down to: " + counter 

		for var i = 1; i <= 3; i++ {
			print "Looping: " + i
		}

		var score = 10
		score += 5
		print "Score after += is: " + score 

		score *= 2
		print "Score after *= is: " + score 

		for var i = 0; i <= 6; i += 2 {
			print "Step: " + i
		}

		var exp = 2
		exp **= 3
		print "Exp after **= is: " + exp 

		var mod = 10
		mod %= 3
		print "Mod after %= is: " + mod 

		var treats = 5
		var hungry = true

		if hungry and treats > 0 {
			print "Time for snacks"
		}

		var tired = false
		if tired or not hungry {
			print "No snacks "
		} else {
			print "Still need food"
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
