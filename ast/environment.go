package ast

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
	_, exists := e.vars[name]
	if exists {
		e.vars[name] = value
		return true
	}
	if e.parent != nil {
		return e.parent.Assign(name, value)
	}
	return false
}
