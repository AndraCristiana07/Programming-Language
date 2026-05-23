package testing

import (
	"bytes"
	"io"
	"my_language/my_language"
	"os"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedVar    any
		expectedOutput string
	}{
		{
			name: "Simple print statement",
			input: `
                var testResult = 0
                print testResult
            `,
			expectedVar:    0,
			expectedOutput: "0",
		},
		{
			name: "Print with addition",
			input: `
                var testResult = 0
                print testResult + 5
            `,
			expectedVar:    0,
			expectedOutput: "5",
		},
		{
			name: "Print with multiplication",
			input: `
				var testResult = 0
				print testResult * 5
			`,
			expectedVar:    0,
			expectedOutput: "0",
		},
		{
			name: "Print with variable update",
			input: `
				var testResult = 0
				testResult = testResult + 5
				print testResult
			`,
			expectedVar:    5,
			expectedOutput: "5",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			tokens := my_language.Lex(tc.input)
			parser := my_language.NewParser(tokens)
			program := parser.Parse()
			env := my_language.NewEnvironment()

			my_language.Eval(program, env)

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			_, _ = io.Copy(&buf, r)
			capturedText := strings.TrimSpace(buf.String())

			result, ok := env.Lookup("testResult")
			if !ok {
				t.Fatalf("Variable 'testResult' was missing from environment state entirely")
			}
			if result != tc.expectedVar {
				t.Errorf("Expected 'testResult' to be %v, got %v", tc.expectedVar, result)
			}

			if capturedText != tc.expectedOutput {
				t.Errorf("Expected print output to be %q, but got %q", tc.expectedOutput, capturedText)
			}
		})
	}
}
