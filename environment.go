package main

type Environment struct {
	vars map[string]int
}

func NewEnvironment() *Environment {
	return &Environment{
		vars: make(map[string]int),
	}
}
