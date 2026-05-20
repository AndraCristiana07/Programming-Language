package main

type Environment struct {
	vars map[string]any
}

func NewEnvironment() *Environment {
	return &Environment{
		vars: make(map[string]any),
	}
}
