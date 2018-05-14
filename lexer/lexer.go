package lexer

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
	charCount := len(ts.Code)
	i := 0
	for i < charCount {

	}

	/*
		tokens = append(tokens, Token{
			TokenType: EOF,
			Span:      nil,
		})
	*/
	ts.Tokens = tokens
}

func (ts *TokenStream) CurrentToken() Token {
	return ts.Tokens[ts.Pos]
}

func (ts *TokenStream) NextToken() Token {
	ts.Pos += 1
	for {
		if ts.Tokens[ts.Pos].TokenType == Comment {
			ts.Pos += 1
		} else {
			break
		}
	}
	return ts.Tokens[ts.Pos]
}

func (ts *TokenStream) isKeyword(value string) bool {
	return value == "main" || value == "number" || value == "string" || value == "bool" || value == "return"
}

func (ts *TokenStream) unexpectedToken(c rune, i int) {
	var lineCount = 1
	var column = 0
	var isFirstLine = true
	for currIndex := i - 1; currIndex > 0; currIndex-- {
		fmt.Println(currIndex)
		if ts.nthChar(currIndex) == "\n" {
			if isFirstLine {
				column = i - currIndex
				isFirstLine = false
			}
			lineCount++
		}
	}
	// panicでいいのか？
	panic(column)
}

func (ts *TokenStream) nthChar(i int) string {
	return string([]rune(ts.Code)[i])
}