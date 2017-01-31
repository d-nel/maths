package main

import (
	"fmt"
	"github.com/d-nel/maths/parse"
)

func main() {
	t := parse.NewTokenizer("-5(12 + 0.34)")

	tok := t.NextToken()
	for tok.Type != parse.TokenNil {
		fmt.Println(tok)
		tok = t.NextToken()
	}
}
