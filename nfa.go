package main

import (
	"fmt"
)

//edge 1 and edge2 are two arrows that come from state
//Pointers to the other states
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

//keeps track of initial & accept state of my fragment of nfa
type nfa struct {
	initial *state
	accept  *state
}

//postfix regular expressions ro nfa

func poregtonfa(pofix string) *nfa {
	//create nfda stack
	//going to be an array of pointers to nfa
	//Give one thats empty
	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}

			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
