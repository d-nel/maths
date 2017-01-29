package main

import (
	"fmt"
	"github.com/d-nel/maths/parse"
)

func main() {
	tok := parse.Token{parse.TokenInteger, "11"}
	fmt.Println(tok)

	t := parse.NewTokenizer("1110+ 12")
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
}
