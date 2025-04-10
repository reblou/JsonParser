package main

import "github.com/golang-collections/collections/stack"

type ParseState int

const (
	ObjectState ParseState = iota
	ArrayState
	StringState
)

func isValidJson(s string) bool {
	// Build stack of states once one state concluded, pop off stack
	state := stack.New()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '{':
			state.Push(ObjectState)
		case '[':
			state.Push(ArrayState)
		case '"':
			state.Push(StringState)
		}
	}
}
