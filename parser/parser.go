package parser

import (
	"github.com/kamakuni/go-toy-interpreter/lexer"
)

type Parser struct {
	TokenStream lexer.TokenStream
	Token       lexer.Token
	Span        lexer.Span
}
