package parse

import (
	"unicode"
)

type TokenType uint16

const (
	TokenNumber TokenType = iota
	TokenSymbol
	TokenNil
)

func (tokType TokenType) String() string {
	switch tokType {
	case TokenNumber:
		return "number"
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

type Parser struct {
	toknr *Tokenizer
}

func NewTokenizer(source string) *Tokenizer {
	return &Tokenizer{[]rune(source), []rune(""), 0, Token{}}
}

func (t *Tokenizer) readNumber(hasPoint bool) Token {
	c := t.currentChar()

	for unicode.IsDigit(c) || (!hasPoint && c == '.') {
		hasPoint = hasPoint || c == '.'

		t.storeChar()
		c = t.nextChar()
	}

	t.tok = Token{TokenNumber, t.popBuffer()}
	return t.tok
}

func (t *Tokenizer) NextToken() Token {
	t.skipSpace()
	c := t.currentChar()

	if unicode.IsDigit(c) {
		return t.readNumber(false)
	}

	switch c {
	case '+', '*', '/', '(', ')':
		t.storeChar()
		t.nextChar()
		t.tok = Token{TokenSymbol, t.popBuffer()}
		return t.tok
	case '-', '.':
		hasPoint := c == '.'
		t.storeChar()
		c = t.nextChar()
		if unicode.IsDigit(c) {
			return t.readNumber(hasPoint)
		}

		t.tok = Token{TokenSymbol, t.popBuffer()}
		return t.tok
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

func (p *Parser) NextNode() Node {
	lhs := p.Expr()

	if p.toknr.tok.Type == TokenSymbol && p.toknr.tok.Value == "*" {
		rhs := p.Expr()
		return &NodeMulti{lhs, rhs}
	}

	return &NodeMulti{}
}

func (p *Parser) Expr() Node {
	tok := p.toknr.NextToken()

	if tok.Type == TokenNumber {
		p.toknr.NextToken()
		return &NodeNumberLit{tok.Value}
	}

	return &NodeNumberLit{"-1"}
}
