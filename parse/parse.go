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
	buffer []rune
	index  int
	tok    Token
}

func NewTokenizer(source string) *Tokenizer {
	return &Tokenizer{[]rune(source), []rune(""), 0, Token{}}
}

func (t *Tokenizer) readNumber() Token {
	t.storeChar()
	c := t.currentChar()
	for unicode.IsDigit(c) {
		t.storeChar()
		c = t.nextChar()
	}

	return Token{TokenInteger, t.popBuffer()}
}

func (t *Tokenizer) NextToken() Token {
//	var value []rune
//	tokType := TokenNil

	t.skipSpace()
	c := t.currentChar()

	if unicode.IsDigit(c) {
		return t.readNumber()
	}

	if c == rune('+') {
		t.storeChar()
		t.nextChar()
		return Token{TokenSymbol, t.popBuffer()}
	}


	// @TODO: Better handling of unknown chars
	return Token{TokenNil, "error"}
}

func (t *Tokenizer) skipSpace() {
	c := t.currentChar()
	for unicode.IsSpace(c) {
		c = t.nextChar()
	}
}

func (t *Tokenizer) storeChar() {
	t.buffer = append(t.buffer, t.currentChar())
}

func (t *Tokenizer) popBuffer() string {
	buff := string(t.buffer)
	t.buffer = []rune("")
	return buff
}

func (t *Tokenizer) nextChar() rune {
	t.index += 1
	return t.currentChar()
}

func (t *Tokenizer) currentChar() rune {
	if t.index < len(t.source) {
		return t.source[t.index]
	}

	return rune(0)
}
