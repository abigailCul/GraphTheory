package main

import (
	"fmt"
)

func intPost(infix string) string {
	/*rune: character as its displayed on the screen
	character converts array of runes back to string.
	specials: map special characters into integers. Keep track
	of specoial characters im allowing*/
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	pofix, s := []rune{}, []rune{}

	//loop over infox string and convert to postfix string
	/* loop through infix and first thing range will return
	is the index of character we are currently reading*/
	for _, r := range infix {
		switch {
		/*
			If we see a closing bracket we are going to pop things of stack
			Until we find the open bracket
		*/
		case r == '(':
			s = append(s, r)
		case r == ')':
			//while last element on stack. Take element of top of stack and stick into pofix
			for s[len(s)-1] != '(' {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] //s[:len(s)-1] : get rid of last element on s
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)

		}
	}

	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}

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
			frag.accept.edge2 = &accept

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

/*
	Takes regular expression from postfix and any strings
	return true if regular expressions match string
	otherwise its returns false.
	True if any string a,b or any number of c's including 0 or more,
	false on anything else
*/
//Takes string - Return boolean back true/false value
func pomatch(po string, InputStr string) bool {

	//Creave variable ismatch by default if string doesnt match
	ismatch := false

	/*
		variable - function called on regular expressions
		On nfa you can be in any number of states at a given time,
		ndfa you are always in one state
	*/
	ponfa := poregtonfa(po)

	/*
		Create an array of states
		Kepp track of current states.
	*/
	current := []*state{}
	/*
		Everytime you read a character from input String inputStr
		look at the list of current states you are in,
		Any state i can move to along arrow
		and any state you get along e arrow
	*/
	next := []*state{}

	// Pass array and change it in another function - convert to slice
	// passed by default
	current = addState(current[:], ponfa.initial, ponfa.accept)

	/*
		Loop through inputStr character at a time - r -
		Read character from that loop through current array
		c - current state im in
	*/
	for _, r := range InputStr {
		for _, c := range current {
			/*
				Current is the same as I'm currently reading from
				from inputStr. Have to ass that state along c arrows
				and add that to next
			*/
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}

		/*
			Add c state and any state i can get to from the c state along e arrows
		*/
		current, next = next, []*state{}
	}

	/*
		Loop through current array in states I am currently in
	*/
	for _, c := range current {
		// the state I'm looping through in current array
		//equals to accept state of po nfa
		//set ismatch = true
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}
func addState(l []*state, s *state, a *state) []*state {

	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

func main() {

	var yourInput string
	var regexp string
	fmt.Println("Please enter the string you would like to test: ")
	fmt.Scan(&yourInput)

	fmt.Println("Please enter the Regex You want to Test the String Against: ")
	fmt.Scan(&regexp)

	fmt.Println(pomatch(regexp, yourInput))

	//following 4 examples of output are testing the 3 basic operations | , * , .
	/*fmt.Println("Infix:      ", "a.b.c*")
	fmt.Println("postFix:    ", intPost("a.b.c*"))

	fmt.Println("Infix:      ", "(a.(b|d))*")
	fmt.Println("postFix:    ", intPost("(a.(b|d))*"))

	fmt.Println("Infix:      ", "a.(b|d).c*")
	fmt.Println("postFix:    ", intPost("a.(b|d).c*"))

	fmt.Println("Infix:      ", "a.(b.b)+.c")
	fmt.Println("postFix:    ", intPost("a.(b.b)+.c"))

	nfa := poregtonfa("ab.c*|")

	// print out what is returned the nfa struct
	fmt.Println("postFix:      ", "a.(b.b)+.c")
	fmt.Println("nfa:         ", nfa)

	fmt.Println(pomatch("ab.c*|", "ccc")) //return true
	//fmt.Println(pomatch("ab.c*|", "def")) //return false*/

}
