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

	//Loops through post fix regular expression a character at a time
	//Switch - pop fragments - pointers to nfa fragments
	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			//append a new pointer to nfa struct and give address of instance
			//push fragment to nfa stack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

			//pop 2 fragments off
			// Push fragment to stack
			//Creates two new states accept and initial
			//join two states to fragments popped of stack
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

		//pop 1 fragment of nfa stack
		//* only works on one fragment of nfa
		case '*':
			//push new fragment
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			//Create new accept and initial
			accept := state{}

			//initial has edge1 as initial state of frag popped off
			//edge 2 needs to point at new accept
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

			//Create new accept state an dinitial state
			//Set symbol to r
			//edge one needs to point to initial - accept point to accept
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
