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

		for (var j = 0; j < 4; j = j + 1) {
			print "I am number: " + j
		}

		for (var k = 0; k < 4; k++) {
			print "I am number: " + k
		}

		var m = 2
		m++
		print "m is: " + m

		var n = 5 ** 2
		print "n is: " + n

		var o = 10 % 3
		print "o is: " + o

		var p = 3
		p **= 3
		print "p is: " + p
	    
		p += 2
		print "p after += 2 is: " + p
		
		p *= 2
		print "p after *= 2 is: " + p

		p /= 2
		print "p after /= 2 is: " + p

		p -= 1
		print "p after -= 1 is: " + p

		p %= 3
		print "p after %= 3 is: " + p
	
		var str1 = "Hello"
		str1 += " World"
		print "String 1: " + str1

		if (5 > 3 and 2 < 4) {
			print "Both conditions are true"
		}

		if (5 > 3 or 2 > 4) {
			print "At least one condition is true"
		}

		if (not (5 < 3)) {
			print "5 is not less than 3"
		}
		
		if ((true or false) and false) {
			print "This won't print"
		} else {
			print "This will print"
		}
	
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
