package parse

import (
	"unicode"
)

type TokenType uint16

const (
	TokenInteger TokenType = iota
	TokenSymbol
	TokenNil
)

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

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	source []rune
	index  int
	tok    Token
}

func NewTokenizer(source string) *Tokenizer {
	return &Tokenizer{[]rune(source), 0, Token{}}
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

	// @Bug: If there is no space between digits and plus it'll be combined
	// into a TokenSymbol.
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

	return rune(0)
}
