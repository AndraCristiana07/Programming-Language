package main

import (
	"bufio"
	"fmt"
	"my_language/ast"
	"my_language/parser"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type PanicErrorListener struct {
	*antlr.DefaultErrorListener
}

func (p *PanicErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("SyntaxError: Line %d:%d - %s", line, column, msg))
}

func StartREPL() {
	eval := ast.NewVisitor()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the language REPL!")
	fmt.Println("Type your commands below. Type 'exit' to quit.")
	fmt.Println("-------------------------------------------")

	for {
		fmt.Print(">>> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if trimmed == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		if trimmed == "" {
			continue
		}

		// protect the shell session from dying on errors
		func() {
			defer func() {
				if r := recover(); r != nil {
					// map based error structure
					if errObjPtr, ok := r.(*map[string]any); ok && errObjPtr != nil {
						errObj := *errObjPtr
						if errorText, exists := errObj["text"].(string); exists {
							fmt.Fprintf(os.Stderr, "Runtime Error: %s\n", errorText)
						} else {
							fmt.Fprintf(os.Stderr, "Runtime Error [%v]: %v (Line %v)\n", errObj["type"], errObj["message"], errObj["line"])
						}
					} else {
						fmt.Fprintf(os.Stderr, "Internal Crash: %v\n", r)
					}
				}
			}()

			// parse the line typed by the user
			input := antlr.NewInputStream(line)
			lexer := parser.NewGrammarLexer(input)
			tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewGrammarParser(tokens)

			// catch syntax typos
			panicListener := &PanicErrorListener{}
			lexer.AddErrorListener(panicListener)
			p.AddErrorListener(panicListener)

			tree := p.Program()

			// execute statement
			tree.Accept(eval)
		}()
	}
}

func runFile(filename string) {

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		os.Exit(1)
	}

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

	inputCode := string(fileBytes)
	input := antlr.NewInputStream(inputCode)

	lexer := parser.NewGrammarLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewGrammarParser(tokens)

	panicListener := &PanicErrorListener{}
	lexer.AddErrorListener(panicListener)
	p.AddErrorListener(panicListener)

	tree := p.Program()

	fmt.Println("--- Parse Tree ---")
	fmt.Println(tree.ToStringTree(nil, p))

	eval := ast.NewVisitor()
	tree.Accept(eval)

	env := eval.GetEnvironment()

	mainFunc, found := env.Lookup("main")
	if found {
		callable, ok := mainFunc.(ast.Callable)
		if !ok {
			panic("TypeError: 'main' is defined but it is not a callable function")
		}

		callable.Call(eval, []any{})
	} else {
		fmt.Println("Warning: No main() function discovered. Executing script top-to-bottom.")
	}

}

func main() {
	if len(os.Args) < 2 {
		StartREPL()
		return
	}

	runFile(os.Args[1])
}
