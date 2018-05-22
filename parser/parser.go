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
