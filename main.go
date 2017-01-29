package main

import (
	"fmt"
	"github.com/d-nel/maths/parse"
)

func main() {
	t := parse.NewTokenizer("-5(12 + .25)")
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
}
