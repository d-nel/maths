package main

import (
	"fmt"
	"github.com/d-nel/maths/parse"
)

func main() {
	/*
		t := parse.NewTokenizer("-5(12 + 0.34) + 52 - .2")

		tok := t.NextToken()
		for tok.Type != parse.TokenNil {
			fmt.Println(tok)
			tok = t.NextToken()
		}
	*/

	parser := parse.NewParser(parse.NewTokenizer("-5 * .2"))

	node := parser.NextNode()
	fmt.Println(node.LHS().Value())
	fmt.Println(node.RHS().Value())

}
