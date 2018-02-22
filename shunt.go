package main

import "fmt"

func intpost(infix string) string {
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

func main() {
	//Answer ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intpost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intpost("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix: ", "(a.(b|d)).c*")
	fmt.Println("Postfix: ", intpost("(a.(b|d)).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix: ", "(a.(b.b))+.c")
	fmt.Println("Postfix: ", intpost("(a.(b.b))+.c"))

}
