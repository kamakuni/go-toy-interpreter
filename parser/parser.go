package parser

import (
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/lexer"
)

type Parser struct {
	TokenStream  lexer.TokenStream
	Token        lexer.Token
	Span         lexer.Span
	TokenCount   int
	CurrentIndex int
}

type RPNValue int

const (
	Operator RPNValue = iota
	Number
)

func CreateParser(tokenStream lexer.TokenStream, span lexer.Span) Parser {
	tokenCount := len(tokenStream.Tokens)
	currentToken := tokenStream.CurrentToken()
	return Parser{
		TokenStream: tokenStream,
		Token:       currentToken,
		TokenCount:  tokenCount,
	}
}

func (p *Parser) CurrnetTokenToString() string {
	return fmt.Sprint(p.Token.TokenType)
}

func (p *Parser) tokenToString(tokenType lexer.TokenType) string {
	return fmt.Sprint(tokenType)
}

func (p *Parser) unexpetedToken(ut string) {
	panic(fmt.Sprintf("Unexpected token found. Expected: %v, Found: %v instead.", ut, p.TokenStream.Tokens[p.CurrentIndex+1].TokenType))
}

func (p *Parser) eatToken(expectedToken string) bool {
	isExist := p.checkToken(expectedToken)

	// If there is a token next, advance token and return true, otherwise return false.
	if isExist {
		return p.advanceToken()
	} else {
		return false
	}
}

func (p *Parser) checkToken(expectedToken string) bool {
	return p.tokenToString(p.TokenStream.Tokens[p.CurrentIndex+1].TokenType) == expectedToken
}

func (p *Parser) advanceToken() bool {
	p.CurrentIndex++

	// If have next token, get next token and return true otherwise return false.
	if p.CurrentIndex != p.TokenCount {
		p.Token = p.TokenStream.NextToken()
	}

	return p.CurrentIndex != p.TokenCount
}
