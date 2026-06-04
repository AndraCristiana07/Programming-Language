package testing

import (
	"fmt"
	"my_language/ast"
	"my_language/parser"
	"sort"

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

func sliceSorter(val any) any {
	if slicePtr, ok := val.(*[]any); ok && slicePtr != nil {
		slice := *slicePtr
		if len(slice) <= 1 {
			return val
		}

		// string slice
		if _, ok := slice[0].(string); ok {
			strSlice := make([]string, len(slice))
			for i, v := range slice {
				strSlice[i] = fmt.Sprintf("%v", v)
			}
			sort.Strings(strSlice)

			sortedSlice := make([]any, len(slice))
			for i, v := range strSlice {
				sortedSlice[i] = v
			}
			return &sortedSlice
		}

		// number slice
		if _, ok := slice[0].(int); ok {
			intSlice := make([]int, len(slice))
			for i, v := range slice {
				if iv, ok := v.(int); ok {
					intSlice[i] = iv
				}
			}
			sort.Ints(intSlice)

			sortedSlice := make([]any, len(slice))
			for i, v := range intSlice {
				sortedSlice[i] = v
			}
			return &sortedSlice
		}
	}
	return val
}
