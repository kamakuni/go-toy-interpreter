package parser

import (
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/ast"
	"github.com/kamakuni/go-toy-interpreter/lexer"
)

type Parser struct {
	TokenStream  lexer.TokenStream
	Token        lexer.Token
	Span         lexer.Span
	TokenCount   int
	CurrentIndex int
}

type RPNValue struct {
	Type  RPNType
	Value interface{}
}

type RPNType int

const (
	Operator RPNType = iota
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

func (p *Parser) unexpectedToken(ut string) {
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

func (p *Parser) eatOperator() bool {
	return p.eatToken("Plus") || p.eatToken("Minus") || p.eatToken("Multiple") || p.eatToken("Divide") || p.eatToken("Mod")
}

func (p *Parser) getCurrentNumber() float64 {
	switch p.Token.TokenType {
	case lexer.Number:
		return p.Token.Value.(float64)
	default:
		panic("Error while parsing to integer.")
	}
}

/*
func (p *Parser) Parse() ast.Expr {
	var block []ast.Expr

	// Read all tokens and create statements, then push it to the block.
	for p.CurrentIndex < p.TokenCount {

		// Ignore the current token if it is useless.
		if p.Token.TokenType == lexer.Comment {
			p.advanceToken()
			continue
		}

		// Determine the parse type for current or (if not enough) next token.
		var stmt ast.Expr
		switch p.Token.TokenType {
		case lexer.Keyword:
			if p.Token.TokenType == "number" {
				stmt = Expr{
					node: p.ParseInteger(),
				}
			}
			TokenType::Keyword(ref x) if x == "string" => {
				Box::new(Expr {
					span: None,
					node: self.parse_string(),
				})
			}
			TokenType::Keyword(ref x) if x == "bool" => {
				Box::new(Expr {
					span: None,
					node: self.parse_bool(),
				})
			}
			TokenType::Identifier(ref x) if x == "if" => {
				Box::new(Expr {
					span: None,
					node: self.parse_if(),
				})
			}
			TokenType::Identifier(ref x) => {
				// Eat LParen
				if self.eat_token("LParen") {
					Box::new(Expr {
						span: None,
						node: self.parse_call(x.clone()),
					})
				} else {
					self.unexpected_token("LParen");
					unimplemented!();
				}
			}
			TokenType::RBrace => break,
			TokenType::EOF => {
				block.push(Box::new(Expr {
					span: None,
					node: Expr_::EOF,
				}));
				break;
			}
			_ => {
				self.unexpected_token(&self.token_to_string(&self.token.token_type));
				Box::new(Expr {
					span: None,
					node: Expr_::Nil,
				})
			}
		};

	}
}
*/

func (p *Parser) parseInteger() interface{} {
	var identifier = ""
	// Eat Identifier
	if p.eatToken("Identifier") {
		if p.Token.TokenType == lexer.Identifier {
			identifier = p.Token.Value.(string)
		} else {
			panic("not yet implemented")
		}
		// Eat equal symbol (=)
		if p.eatToken("Equals") {
			return p.calculate(identifier)
		} else {
			p.unexpectedToken("Equals")
		}
	} else {
		p.unexpectedToken("Identifier")
	}
	return ast.Nil{}
}

/**
* Calculate arithmetic expression with Shunting-Yard Algorithm
 */
func (p *Parser) calculate(identifier string) interface{} {
	var operatorStack []lexer.TokenType
	var rpn []RPNValue
	var opPrecedences = make(map[ast.TokenType]int)
	var waitExp = true

	// Push operators to precendeces list
	opPrecedences[lexer.Plus] = 2
	opPrecedences[lexer.Minus] = 2
	opPrecedences[lexer.Multiple] = 3
	opPrecedences[lexer.Divide] = 3
	opPrecedences[lexer.Mod] = 3

	// Loop for all numbers and operators
	for {
		if p.EatToken("Number") {
			// Get first number
			rpn = append(rpn, RPNValue{Type: Number, Value: p.getCurrentNumber()})
			waitExp = false
		} else if waitExp {
			// If number is not set break the loop
			break
		} else if p.eatOperator() {
			// If eat an operator
			var stackLen = len(operatorStack)

			for stackLen > 0 &&
				opPrecedences[p.Token.TokenType] <
					opPrecedences[operatorStack[stackLen-1]] {
				rpn = append(rpn, RPNValue{Type: Operator, Value: operatorStack[stackLen-1]})
				operator_stack = append(operator_stack[:stackLen-1], operator_stack[stackLen-1+1:]...)
				stackLen--
			}

			operatorStack = append(operatorStack, p.Token.TokenType)
			waitExp = true
		} else {
			// This means expression is ended and we need a semicolon check.
			p.expectSemicolon()
			break
		}
	}

	// waitExp == true means line ended with an operator or line is empty.
	if waitExp {
		p.unexpectedToken("Number")
	}

	// Popping stack and pushing to rpn queue.
	for i := len(operatorStack) - 1; i >= 0; i++ {
		rpn = append(rpn, RPNValue{
			Type:  Operator,
			Value: operatorStack[i],
		}) //operatorStack[i]
	}

	// Calling soveRPN function and returning it as Expr_.
	return ast.Assign{
		Value: identifier,
		Expr: ast.Expr{
			Node: ast.Constant{
				Type: ast.Number{
					Value: p.solveRpn(rpn),
				},
			},
		},
	}
}

func (p *Parser) solveRpn(rpn []RPNValue) float64 {
	var valStack []float64

	for index, value := range rpn {
		if value.Type == Number {
			valStack = append(valStack, value.Value)
		} else if value.Type == Operator {
			stackLength := len(valStack)
			if stackLength >= 2 {
				first, valStack = pop(valStack)
				second, valStack = pop(valStack)
				switch value.Value {
				case lexer.Plus:
					valStack = append(valStack, second+first)
				case lexer.Minus:
					valStack = append(valStack, second-first)
				case lexer.Multiple:
					valStack = append(valStack, second*first)
				case lexer.Divide:
					valStack = append(valStack, second/first)
				case lexer.Mod:
					valStack = append(valStack, second%first)
				default:
					p.unexpectedToken(p.tokenToString(x))
				}
			} else {
				panic("Parse error in arithmetic value. Check number assignment.")
			}
		}
	}
	return valStack[0]
}

func (p *Parser) parseString() {

}

func pop(slice []interface{}) (interface{}, []interface{}) {
	ans := src[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans, slice
}

func (p *Parser) expectedSemicolon() {
	if p.eatToken("Semicolon") {
		p.advanceToken()
	} else {
		p.unexpectedToken("Semicolon")
	}
}
