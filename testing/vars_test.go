package testing

import (
	"bytes"
	"io"
	"my_language/ast"
	"os"
	"strings"
	"testing"
)

func TestUnusedVariableWarnings(t *testing.T) {
	env := ast.NewEnvironment()

	env.DefineWithLine("ghostVar", 99, 3)

	env.DefineWithLine("alpha", true, 5)
	env.DefineWithLine("beta", false, 5)

	_, _ = env.Lookup("alpha")

	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	env.CheckUnusedVars()

	w.Close()
	os.Stderr = oldStderr

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Variable 'ghostVar' declared on line 3 is unused") {
		t.Errorf("Expected warning for 'ghostVar' on line 3, got:\n%s", output)
	}

	if !strings.Contains(output, "Variable 'beta' declared on line 5 is unused") {
		t.Errorf("Expected warning for 'beta' on line 5, got:\n%s", output)
	}

}
