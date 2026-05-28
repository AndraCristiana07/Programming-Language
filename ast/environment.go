package ast

import "fmt"

type Environment struct {
	vars   map[string]any
	parent *Environment
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
		if _, isNative := existingValue.(*NativeFunction); isNative {
			panic(fmt.Sprintf("CompileError: '%s' is a protected built-in function and cannot be redefined", name))
		}
	}
	e.vars[name] = value
}

// go to parent until finds the variable
func (e *Environment) Lookup(name string) (any, bool) {
	val, exists := e.vars[name]
	if exists {
		return val, true
	}
	if e.parent != nil {
		return e.parent.Lookup(name)
	}
	return nil, false
}

// go up the scopes to update a variable
func (e *Environment) Assign(name string, value any) bool {
	if existingValue, exists := e.vars[name]; exists {
		if _, isNative := existingValue.(*NativeFunction); isNative {
			panic(fmt.Sprintf("CompileError: Cannot assign a value to protected built-in function '%s'", name))
		}

		e.vars[name] = value
		return true
	}

	if e.parent != nil {
		return e.parent.Assign(name, value)
	}

	return false
}
