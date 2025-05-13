package main

import "github.com/golang-collections/collections/stack"

type ParseState int

const (
	ObjectState ParseState = iota
	ArrayState
	StringState
)

var strategy func(*stack.Stack, byte)

func isValidJson(s string) bool {
	// Build stack of states once one state concluded, pop off stack
	state := stack.New()
	strategy = NormalCharParse
	for i := 0; i < len(s); i++ {
		strategy(state, s[i])
	}

	if state.Len() > 0 {
		panic("Incomplete Json")
	}
	return true
}

func NormalCharParse(state *stack.Stack, c byte) {
	switch c {
	case '{':
		state.Push(ObjectState)
	case '[':
		state.Push(ArrayState)
	case '"':
		strategy = StringCharParse
	case '}':
		if l := state.Pop(); l != ObjectState {
			panic("Invalid Json, unexpected }")
		}
	case ']':
		if l := state.Pop(); l != ArrayState {
			panic("Invalid Json, unexpected ]")
		}
	}
}

func StringCharParse(state *stack.Stack, c byte) {
	if c == '"' {
		// end string parse
		strategy = NormalCharParse
	}
}
