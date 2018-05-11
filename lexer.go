package main

import (
	"fmt"
)

type TokenType int

const (
	Keyword      TokenType = iota // like int string or let
	Identifier                    // like variable names
	Char                          // Char variables inside " ' "
	String                        // String variables inside quotes
	Number                        // Number variable
	True                          // Boolean true
	False                         // Boolean false
	Equals                        // =
	Plus                          // +
	Minus                         // -
	Multiple                      // *
	Divide                        // /
	Mod                           // %
	Greater                       // >
	Lesser                        // <
	GreaterEqual                  // >=
	LesserEqual                   // <=
	LParen                        // (
	RParen                        // )
	LBrace                        // {
	RBrace                        // }
	LBracket                      // [
	RBracket                      // ]
	Comma                         //
	Semicolon                     // ;
	Comment                       // '//'
	EOF                           // End of File
)

type Span struct {
	Lo int
	Hi int
}

type Token struct {
	TokenType TokenType
	Span      Span
}

type TokenStream struct {
	Code   string
	Tokens []Token
	Pos    int
	Curr   string
}

func CreateTokenStream(code string) (ts TokenStream) {
	ts = TokenStream{
		Code:   code,
		Tokens: []Token{},
		Pos:    0,
		Curr:   "",
	}
	ts.Tokenize()
	return
}

func (ts *TokenStream) Tokenize() {
	// Todo create tokenize function
	tokens := []Token{}
	char_count := len(ts.Code)
	i := 0
	for {
		fmt.Println(tokens)
		fmt.Println(char_count)
		fmt.Println(i)
	}
}

func (ts *TokenStream) CurrentToken() Token {
	return ts.Tokens[ts.Pos]
}
