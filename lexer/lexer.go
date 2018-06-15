package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

// like int, string or let
type Keyword struct {
	Value: string
}
// like variable names
type Identifier struct {
	Value: string
}
// Char variables, inside " ' "
type Char struct {
	Value: string
}
// String variables, inside quotes
type String struct {
	Value: string
}
// Number variable
type Number struct {
	Value: string
}

type TokenType int

const (
/*	Keyword      TokenType = iota // like int string or let
	Identifier                    // like variable names
	Char                          // Char variables inside " ' "
	String                        // String variables inside quotes
	Number                        // Number variable*/
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

func (t TokenType) String() string {
	switch t {
	case Keyword:
		return "Keyword"
	case Identifier:
		return "Identifier"
	case Char:
		return "Char"
	case String:
		return "String"
	case Number:
		return "Number"
	case True:
		return "True"
	case False:
		return "False"
	case Equals:
		return "Equals"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Multiple:
		return "Multiple"
	case Divide:
		return "Divide"
	case Mod:
		return "Mod"
	case Greater:
		return "Greater"
	case Lesser:
		return "Lesser"
	case GreaterEqual:
		return "GreaterEqual"
	case LesserEqual:
		return "LesserEqual"
	case LParen:
		return "LParen"
	case RParen:
		return "RParen"
	case LBrace:
		return "LBrace"
	case RBrace:
		return "RBrace"
	case LBracket:
		return "LBracket"
	case RBracket:
		return "RBracket"
	case Comma:
		return "Comma"
	case Semicolon:
		return "Semicolon"
	case Comment:
		return "Comment"
	case EOF:
		return "EOF"
	default:
		return "OTHER TYPE"
	}
}

type Span struct {
	Lo int
	Hi int
}

type Token struct {
	TokenType TokenType
	Span      *Span
	Value     interface{}
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
	var i = 0
	for i < charCount {
		currentChar := ts.nthChar(i)

		if unicode.IsSpace(currentChar) {
			i++
		} else if unicode.IsLetter(currentChar) {
			var tmp = ""
			for i < charCount && unicode.IsLetter(ts.nthChar(i)) {
				tmp = tmp + string(ts.nthChar(i))
				i++
			}
			var tmpStr = strings.ToLower(tmp)
			if ts.isKeyword(tmpStr) {
				tokens = append(tokens, Token{
					TokenType: Keyword,
					Span:      nil,
					Value:     tmpStr,
				})
			} else if string(tmpStr) == "true" {
				tokens = append(tokens, Token{
					TokenType: True,
					Span:      nil,
				})
			} else if string(tmpStr) == "false" {
				tokens = append(tokens, Token{
					TokenType: False,
					Span:      nil,
				})
			} else {
				tokens = append(tokens, Token{
					TokenType: Identifier,
					Span:      nil,
					Value:     tmpStr,
				})
			}

		} else if unicode.IsNumber(currentChar) {
			var tmp = ""
			for i < charCount && unicode.IsNumber(ts.nthChar(i)) {
				tmp = tmp + string(ts.nthChar(i))
				i++
			}
			tokens = append(tokens, Token{
				TokenType: Number,
				Span:      nil,
				Value:     tmp,
			})
			// If current char is a starting of a string
		} else if currentChar == '"' {
			var tmp = ""
			i++

			for i < charCount && ts.nthChar(i) != '"' {
				tmp = tmp + string(ts.nthChar(i))
				i++
			}

			i++
			tokens = append(tokens, Token{
				TokenType: String,
				Span:      nil,
				Value:     tmp,
			})
			// If current char is a real char
		} else if currentChar == '\'' {
			var tmp = string([]rune(ts.Code)[i+1])
			i = i + 2
			if ts.nthChar(i) == '\'' {
				tokens = append(tokens, Token{
					TokenType: Char,
					Span:      nil,
					Value:     tmp,
				})
				i++
			} else {
				ts.unexpectedToken(ts.nthChar(i), i)
			}
			// If current char is an equals (=)
		} else if currentChar == '=' {
			tokens = append(tokens, Token{
				TokenType: Equals,
				Span:      nil,
			})
			i++
			// If current char is a plus (+)
		} else if currentChar == '+' {
			tokens = append(tokens, Token{
				TokenType: Plus,
				Span:      nil,
			})
			i++
			// If current char is a minus (-)
		} else if currentChar == '-' {
			tokens = append(tokens, Token{
				TokenType: Minus,
				Span:      nil,
			})
			i++
			// If current char is a multiple (*)
		} else if currentChar == '*' {
			tokens = append(tokens, Token{
				TokenType: Multiple,
			})
			i++
			// If current char is a divide (/) or comment (start with //)
		} else if currentChar == '/' {
			i++
			if i < charCount && ts.nthChar(i) == '/' {
				for i < charCount && ts.nthChar(i) == '\n' {
					i++
				}

				i++
				tokens = append(tokens, Token{
					TokenType: Comment,
				})
			} else {
				tokens = append(tokens, Token{
					TokenType: Divide,
				})
			}
			// If currnet char is a mod (%)
		} else if currentChar == '%' {
			tokens = append(tokens, Token{
				TokenType: Mod,
			})
			i++
			// If current char is a greater than (>) or greater than or equal to (>=)
		} else if currentChar == '>' {
			if i+1 < charCount && []rune(ts.Code)[i+1] == '=' {
				tokens = append(tokens, Token{
					TokenType: GreaterEqual,
				})
				i++
			} else {
				tokens = append(tokens, Token{
					TokenType: Greater,
				})
			}
			i++
			// If current char is a lesser than (<) or lesser than or equal to (<=)
		} else if currentChar == '<' {
			if i+1 < charCount && []rune(ts.Code)[i+1] == '=' {
				tokens = append(tokens, Token{
					TokenType: LesserEqual,
				})
				i++
			} else {
				tokens = append(tokens, Token{
					TokenType: Lesser,
				})
			}
			i++
			// If current char is an Open Paranthesis ( ( )
		} else if currentChar == '(' {
			tokens = append(tokens, Token{
				TokenType: LParen,
			})
			i++
			// If current char is a Close Paranthesis ( ) )
		} else if currentChar == ')' {
			tokens = append(tokens, Token{
				TokenType: RParen,
			})
			i++
			// If current char is an Open Braces ( { )
		} else if currentChar == '{' {
			tokens = append(tokens, Token{
				TokenType: LBrace,
			})
			i++
			// If current char is a Close Braces ( } )
		} else if currentChar == '}' {
			tokens = append(tokens, Token{
				TokenType: RBrace,
			})
			i++
			// If current char is an Open Brackets ( [ )
		} else if currentChar == '[' {
			tokens = append(tokens, Token{
				TokenType: LBracket,
			})
			i++
			// If current char is a Close Brackets ( [ )
		} else if currentChar == ']' {
			tokens = append(tokens, Token{
				TokenType: RBracket,
			})
			i++
			// If current char is a Comma ( , )
		} else if currentChar == ',' {
			tokens = append(tokens, Token{
				TokenType: Comma,
			})
			i++
			// If current char is a semicolon ( ; )
		} else if currentChar == ';' {
			tokens = append(tokens, Token{
				TokenType: Semicolon,
			})
			// Else throw an exception
		} else {
			ts.unexpectedToken(currentChar, i)
		}
	}
	// End od file Token
	tokens = append(tokens, Token{
		TokenType: EOF,
	})

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
		if string(ts.nthChar(currIndex)) == "\n" {
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

func (ts *TokenStream) nthChar(i int) rune {
	return []rune(ts.Code)[i]
}
