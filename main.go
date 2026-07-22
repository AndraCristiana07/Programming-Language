package main

import (
	"fmt"
	"io"
	"my_language/ast"
	"my_language/parser"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/chzyer/readline"
)

type PanicErrorListener struct {
	*antlr.DefaultErrorListener
}

func (p *PanicErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("SyntaxError: Line %d:%d - %s", line, column, msg))
}

func parseFileToTree(filename string) any {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("ImportError: Failed to read module '%s': %v", filename, err))
	}

	input := antlr.NewInputStream(string(fileBytes))
	lexer := parser.NewGrammarLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGrammarParser(tokens)

	panicListener := &PanicErrorListener{}
	lexer.AddErrorListener(panicListener)
	p.AddErrorListener(panicListener)

	return p.Program() // returns the parsed *parser.ProgramContext
}

func StartREPL() {
	eval := ast.NewVisitor()

	eval.ImportHandler = parseFileToTree
	// init the Readline instance
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          ">>> ",
		HistoryFile:     ".repl_history", // saves command history to a local dotfile
		InterruptPrompt: "^C",            // prints ^C  on manual cancels
		EOFPrompt:       "exit",          // matches standard termination keyword
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize terminal engine: %v\n", err)
		return
	}
	defer rl.Close() // clean terminal

	fmt.Println("Welcome to the language REPL!")
	fmt.Println("Type your commands below. Type 'exit' to quit.")
	fmt.Println("-------------------------------------------")

	var blockBuilder strings.Builder
	openBraces := 0

	for {
		// read a line from the text engine
		line, err := rl.Readline()

		// handle special keyboard inputs
		if err == readline.ErrInterrupt {
			// User pressed Ctrl+C
			if openBraces > 0 {
				// clear any partially typed multi-line blocks
				openBraces = 0
				blockBuilder.Reset()
				rl.SetPrompt(">>> ")
				fmt.Println(" (Cleared block)")
				continue
			}
			break // quit REPL
		} else if err == io.EOF {
			// user pressed Ctrl+D
			fmt.Println("Goodbye!")
			break
		}

		trimmed := strings.TrimSpace(line)

		if openBraces == 0 && trimmed == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		if trimmed == "" && openBraces == 0 {
			continue
		}

		// track braces on line
		// count open and close braces
		openBraces += strings.Count(line, "{") - strings.Count(line, "}")
		if openBraces < 0 {
			openBraces = 0
		}

		blockBuilder.WriteString(line)
		blockBuilder.WriteString("\n")

		// choose prompt based on whether inside a multi-line block
		if openBraces > 0 {
			rl.SetPrompt("... ")
			continue
		} else {
			rl.SetPrompt(">>> ")
		}

		// pull entire completed block and reset tracking variables
		sourceCode := blockBuilder.String()
		blockBuilder.Reset()

		// protect the shell session from dying on errors
		func() {
			defer func() {
				if r := recover(); r != nil {
					if errObjPtr, ok := r.(*map[string]any); ok && errObjPtr != nil {
						errObj := *errObjPtr
						if errorText, exists := errObj["text"].(string); exists {
							fmt.Fprintf(os.Stderr, "Runtime Error: %s\n", errorText)
						} else {
							fmt.Fprintf(os.Stderr, "Runtime Error [%v]: %v (Line %v)\n", errObj["type"], errObj["message"], errObj["line"])
						}
					} else if syntaxErr, isString := r.(string); isString {
						fmt.Fprintf(os.Stderr, "%s\n", syntaxErr)
					} else {
						fmt.Fprintf(os.Stderr, "Internal Crash: %v\n", r)
					}
				}
			}()

			input := antlr.NewInputStream(sourceCode)
			lexer := parser.NewGrammarLexer(input)
			tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewGrammarParser(tokens)

			panicListener := &PanicErrorListener{}
			lexer.AddErrorListener(panicListener)
			p.AddErrorListener(panicListener)

			tree := p.Program()
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

	eval.ImportHandler = parseFileToTree

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
