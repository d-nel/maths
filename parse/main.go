package main

import (
	"fmt"
	"unicode"
)

type TokenType uint16

const (
 	TokenInteger TokenType = iota
	TokenSymbol
	TokenNil
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	source []rune
	index  int
	tok    Token
}

func (t *Tokenizer) NextToken() Token {
	var value []rune
	tokType := TokenNil

	t.SkipSpace()
	c := t.CurrentChar()

	for unicode.IsDigit(c) {
		value = append(value, c)
		c = t.NextChar()
		tokType = TokenInteger
	}

	if c == rune('+') {
		value = append(value, c)
		c = t.NextChar()
		tokType = TokenSymbol
	}

	return Token{tokType, string(value)}
}

func (t *Tokenizer) SkipSpace() {
	c := t.CurrentChar()
	for unicode.IsSpace(c) {
		c = t.NextChar()
	}
}

func (t *Tokenizer) NextChar() rune {
	t.index += 1
	return t.CurrentChar()
}

func (t *Tokenizer) CurrentChar() rune {
	if t.index < len(t.source) {
		return t.source[t.index]
	}

	return rune('F') // TODO: return blank rune
}


func main() {
	tok := Token{TokenInteger, "11"}
	fmt.Println(tok)

	t := Tokenizer{[]rune("1110 + 12"), 0, tok}
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
	fmt.Println(t.NextToken())
}


func (tokType TokenType) String() string {
	switch tokType {
	case TokenInteger:
		return "int"
	case TokenSymbol:
		return "symbol"
	default:
		return "unknown"
	}
}
