package main

import (
	"fmt"
	"my_language/ast"
	"my_language/parser"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	inputCode := string(fileBytes)
	input := antlr.NewInputStream(inputCode)

	lexer := parser.NewGrammarLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGrammarParser(tokens)

	tree := p.Program()

	fmt.Println("--- Parse Tree ---")
	fmt.Println(tree.ToStringTree(nil, p))

	defer func() {
		if r := recover(); r != nil {
			// catch custom error
			if errObjPtr, ok := r.(*map[string]any); ok && errObjPtr != nil {
				errObj := *errObjPtr

				if errorText, exists := errObj["text"].(string); exists {
					fmt.Fprintf(os.Stderr, "\n Runtime Panic: %s\n", errorText)
				} else {
					fmt.Fprintf(os.Stderr, "\n Runtime Panic [%v]: %v (Line %v)\n", errObj["type"], errObj["message"], errObj["line"])
				}
				os.Exit(1)
			}

			// fallback to standard Go compiler bugs
			fmt.Fprintf(os.Stderr, "\n Internal Interpreter Crash: %v\n", r)
			os.Exit(1)
		}
	}()

	eval := ast.NewVisitor()
	tree.Accept(eval)

}
