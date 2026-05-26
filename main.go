package main

import (
	"fmt"
	"my_language/ast"
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

		var bitand = 5 & 3
		print "5 & 3 is: " + bitand

		var bitor = 5 | 3
		print "5 | 3 is: " + bitor

		var bitxor = 5 ^ 3
		print "5 ^ 3 is: " + bitxor

		var bitlshift = 2 << 2
		print "2 << 2 is: " + bitlshift

		var bitrshift = 8 >> 2
		print "8 >> 2 is: " + bitrshift

		var bitnot = ~5
		print "~5 is: " + bitnot
	  
		var arr = [1, 2, 3]
		print "Array: " + arr

		print "len arr is: " + len(arr)
		arr = append(arr, 4)
		print "Modified append: " + arr


		print "Array element at index 1: " + arr[1]
		arr[0] = 10
		print "Modified Array: " + arr

		func add(a, b) {
			return a + b
		}
		var sum1 = add(5, 7)
		print "Sum1: " + sum1

		func factorial(n) {
			if (n == 0) {
				return 1
			} else {
				return n * factorial(n - 1)
			}
		}
		var fact5 = factorial(5)
		print "Factorial of 5: " + fact5
		print "Factorial of 0: " + factorial(0)

		var bitcomp = 5
		bitcomp &= 5
		print bitcomp

		var toStr = str(5)
		print toStr
		print "type tostr: " + type(toStr)

		var toInt = int("5")
		print toInt
		print "type toInt: " + type(toInt)

		var testBool = false
		print "type pf testBool: " + type(testBool)
		print "type of arr: " + type(arr)
		print "type of arr[0]: " + type(arr[0])

		var arr1 = [[1,2], [3,5], [6,4]]
		arr1 = append(arr1, [6])

		print "abs of -5: " + abs(-5)
		print "min : " + min(2, 6)
		print "max: " + max(3, 8)

		var toChr = chr(67)
		print toChr

		var toOrd = ord("c")
		print toOrd

		var toBool = bool(20)
		print "to bool: " + toBool

		print "range with 1-10: " + range(1, 10)

		print "range with 9: " + range(9)
		var arr2 = [1, 2, 2, 5, 1, 7]
		print "set from arr: " + set(arr2)

		print lower("HellLo")
		print upper("HeellsO")

		print pop(arr2)
		print arr2
		pop(arr2)
		print arr2

		print reverse(arr2)
		print find(arr2, 2)

		var boolArr = [true, false, false]
		print all(boolArr)
		print any(boolArr)
		print float(5)

		print sum(arr2)
		var phrase = "  hello-there  "
		print "stripped: " + strip(phrase)
		print "split: " + split(phrase, "-")
		var filename = "script.exe"
		print startswtih(filename, "script")          
		print endswitsh(filename, ".txt")
		`

	input := antlr.NewInputStream(inputCode)

	lexer := parser.NewGrammarLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGrammarParser(tokens)

	tree := p.Program()

	fmt.Println("--- Parse Tree ---")
	fmt.Println(tree.ToStringTree(nil, p))

	eval := ast.NewVisitor()
	tree.Accept(eval)

}
