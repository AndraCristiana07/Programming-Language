package ast

import (
	"fmt"
	"os"
)

type Environment struct {
	vars   map[string]any
	parent *Environment
}

type VariableMetadata struct {
	Value      any
	IsUsed     bool
	LineNumber int
	Name       string
}

func NewEnvironment() *Environment {
	return &Environment{
		vars:   make(map[string]any),
		parent: nil,
	}
}

// scope
func NewScope(parent *Environment) *Environment {
	return &Environment{
		vars:   make(map[string]any),
		parent: parent,
	}
}

// sets a var strictly in the curr block
func (e *Environment) Define(name string, value any) {
	if existingValue, exists := e.vars[name]; exists {
		actualVal := existingValue
		if meta, ok := existingValue.(*VariableMetadata); ok {
			actualVal = meta.Value
		}
		if _, isNative := actualVal.(*NativeFunction); isNative {
			panic(fmt.Sprintf("CompileError: '%s' is a protected built-in function and cannot be redefined", name))
		}
	}

	e.vars[name] = value
}

func (e *Environment) DefineWithLine(name string, value any, line int) {
	if existingValue, exists := e.vars[name]; exists {
		actualVal := existingValue
		if meta, ok := existingValue.(*VariableMetadata); ok {
			actualVal = meta.Value
		}
		if _, isNative := actualVal.(*NativeFunction); isNative {
			panic(fmt.Sprintf("CompileError: '%s' is a protected built-in function and cannot be redefined", name))
		}
	}

	e.vars[name] = &VariableMetadata{
		Value:      value,
		IsUsed:     false,
		LineNumber: line,
		Name:       name,
	}
}

// go to parent until finds the variable
func (e *Environment) Lookup(name string) (any, bool) {
	val, exists := e.vars[name]
	if exists {
		// tracked user variable -> mark as read and return the val
		if meta, ok := val.(*VariableMetadata); ok {
			meta.IsUsed = true
			return meta.Value, true
		}
		return val, true
	}
	if e.parent != nil {
		return e.parent.Lookup(name)
	}
	return nil, false
}

// go up the scopes to update a variable
func (e *Environment) Assign(name string, value any) bool {
	existingValue, exists := e.vars[name]
	if exists {
		actualVal := existingValue
		if meta, ok := existingValue.(*VariableMetadata); ok {
			actualVal = meta.Value
		}

		if _, isNative := actualVal.(*NativeFunction); isNative {
			panic(fmt.Sprintf("CompileError: Cannot assign a value to protected built-in function '%s'", name))
		}

		// update the value
		if meta, ok := existingValue.(*VariableMetadata); ok {
			meta.Value = value
		} else {
			e.vars[name] = value
		}
		return true
	}

	if e.parent != nil {
		return e.parent.Assign(name, value)
	}

	return false
}

func (e *Environment) CheckUnusedVars() {
	for name, val := range e.vars {
		if meta, ok := val.(*VariableMetadata); ok {
			if name == "main" {
				continue
			}
			if !meta.IsUsed {
				fmt.Fprintf(os.Stderr, "Warning: Variable '%s' declared on line %d is unused\n", name, meta.LineNumber)
			}
		}
	}
}
