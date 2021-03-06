package parser

import (
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/ast"
	"github.com/kamakuni/go-toy-interpreter/lexer"
	"math"
	"strconv"
)

type Parser struct {
	TokenStream  lexer.TokenStream
	Token        lexer.Token
	Span         lexer.Span
	TokenCount   int
	CurrentIndex int
}

type RPNValue struct {
	Value interface{}
}

/*type RPNType int

const (
	Operator RPNType = iota
	Number
)*/

type Operator struct {
	Value lexer.TokenType
}

type Number struct {
	Value float64
}

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

func (p *Parser) tokenToString(tokenType interface{}) string {
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
	t, ok := p.Token.TokenType.(lexer.Number)
	if !ok {
		panic("Error while parsing to float.")
	}
	f, err := strconv.ParseFloat(t.Value, 64)
	if err != nil {
		panic("Error while parsing to float.")
	}
	return f
}

func (p *Parser) Parse() ast.Expr {
	var block []ast.Expr

FOR_LABEL:
	// Read all tokens and create statements, then push it to the block.
	for p.CurrentIndex < p.TokenCount {

		// Ignore the current token if it is useless.
		if t, ok := p.Token.TokenType.(lexer.TokenType); ok && t == lexer.Comment {
			p.advanceToken()
			continue
		}

		// Determine the parse type for current or (if not enough) next token.
		var stmt ast.Expr
		switch t := p.Token.TokenType.(type) {
		case lexer.Keyword:
			if t.Value == "number" {
				stmt = ast.Expr{
					Node: p.parseInteger(),
				}
			} else if t.Value == "string" {
				stmt = ast.Expr{
					Node: p.parseString(),
				}
			} else if t.Value == "bool" {
				stmt = ast.Expr{
					Node: p.parseBool(),
				}
			}
		case lexer.Identifier:
			if t.Value == "if" {
				stmt = ast.Expr{
					Node: p.parseIf(),
				}
			} else {
				// Eat LParen
				if p.eatToken("LParen") {
					stmt = ast.Expr{
						Node: p.parseCall(t.Value),
					}
				} else {
					p.unexpectedToken("LParen")
					//unimplemented!();
				}
			}
		case lexer.TokenType:
			if t == lexer.RBrace {
				break FOR_LABEL
			} else if t == lexer.EOF {
				block = append(block, ast.Expr{
					Node: ast.EOF{},
				})
				break FOR_LABEL
			}
		default:
			p.unexpectedToken(p.tokenToString(t))
			stmt = ast.Expr{
				Node: ast.Nil{},
			}
		}
		block = append(block, stmt)
	}
	return ast.Expr{
		Node: ast.Block{
			Exprs: block,
		},
	}
}

func (p *Parser) parseInteger() interface{} {
	var identifier = ""
	// Eat Identifier
	if p.eatToken("Identifier") {
		if t, ok := p.Token.TokenType.(lexer.Identifier); ok {
			identifier = t.Value
		} else {
			panic("Error while parsing to integer")
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
	var opPrecedences = make(map[lexer.TokenType]int)
	var waitExp = true

	// Push operators to precendeces list
	opPrecedences[lexer.Plus] = 2
	opPrecedences[lexer.Minus] = 2
	opPrecedences[lexer.Multiple] = 3
	opPrecedences[lexer.Divide] = 3
	opPrecedences[lexer.Mod] = 3

	// Loop for all numbers and operators
	for {
		if p.eatToken("Number") {
			// Get first number
			rpn = append(rpn, RPNValue{Value: Number{Value: p.getCurrentNumber()}})
			waitExp = false
		} else if waitExp {
			// If number is not set break the loop
			break
		} else if p.eatOperator() {
			// If eat an operator
			var stackLen = len(operatorStack)
			tokenType := p.Token.TokenType.(lexer.TokenType)
			for stackLen > 0 &&
				opPrecedences[tokenType] <
					opPrecedences[operatorStack[stackLen-1]] {
				rpn = append(rpn, RPNValue{Value: Operator{Value: operatorStack[stackLen-1]}})
				operatorStack = append(operatorStack[:stackLen-1], operatorStack[stackLen-1+1:]...)
				stackLen--
			}

			operatorStack = append(operatorStack, tokenType)
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
	for i := len(operatorStack) - 1; i >= 0; i-- {
		rpn = append(rpn, RPNValue{
			Value: Operator{
				Value: operatorStack[i],
			},
		}) //operatorStack[i]
	}

	// Calling soveRPN function and returning it as Expr_.
	return ast.Assign{
		Value: identifier,
		Expr: ast.Expr{
			Node: ast.Constant{
				Value: ast.Number{
					Value: p.solveRpn(rpn),
				},
			},
		},
	}
}

func (p *Parser) solveRpn(rpn []RPNValue) float64 {
	var valStack []float64

	for _, value := range rpn {
		switch v := value.Value.(type) {
		case Number:
			valStack = append(valStack, v.Value)
		case Operator:
			stackLength := len(valStack)
			if stackLength >= 2 {
				first, valStack := pop(valStack)
				second, valStack := pop(valStack)
				switch v.Value {
				case lexer.Plus:
					valStack = append(valStack, second+first)
				case lexer.Minus:
					valStack = append(valStack, second-first)
				case lexer.Multiple:
					valStack = append(valStack, second*first)
				case lexer.Divide:
					valStack = append(valStack, second/first)
				case lexer.Mod:
					valStack = append(valStack, math.Mod(second, first))
				default:
					p.unexpectedToken(p.tokenToString(v.Value))
				}
			} else {
				panic("Parse error in arithmetic value. Check number assignment.")
			}
		}
	}
	return valStack[0]
}

func (p *Parser) parseString() interface{} {
	var identifier = ""
	var str = ""
	var expr interface{}

	// Eat identifier
	if p.eatToken("Identifier") {
		if i, ok := p.Token.TokenType.(lexer.Identifier); ok {
			identifier = i.Value
		} else {
			panic("not yet implemented")
		}

		if p.eatToken("Equals") {
			if p.eatToken("String") {
				if s, ok := p.Token.TokenType.(lexer.String); ok {
					str = s.Value

					expr = ast.Assign{
						Value: identifier,
						Expr: ast.Expr{
							Node: ast.Constant{
								Value: ast.String{
									Value: str,
								},
							},
						},
					}

					p.expectSemicolon()
					return expr
				} else {
					panic("not yet implemented")
				}
			}
		} else {
			p.unexpectedToken("Equals")
		}
	} else {
		p.unexpectedToken("Identifier")
	}
	return ast.Nil{}
}

func (p *Parser) parseBool() interface{} {
	var identifier = ""
	var boolVal bool
	var expr interface{}

	// Eat identifier
	if p.eatToken("Identifier") {
		if i, ok := p.Token.TokenType.(lexer.Identifier); ok {
			identifier = i.Value
		} else {
			panic("not yet implemented")
		}

		if p.eatToken("Equals") {
			if p.eatToken("True") || p.eatToken("False") {
				if t, ok := p.Token.TokenType.(lexer.TokenType); ok && t == lexer.True {
					boolVal = true
				} else if f, ok := p.Token.TokenType.(lexer.TokenType); ok && f == lexer.False {
					boolVal = false
				} else {
					panic("not yet implemented")
				}
				expr = ast.Assign{
					Value: identifier,
					Expr: ast.Expr{
						Node: ast.Constant{
							Value: ast.Bool{
								Value: boolVal,
							},
						},
					},
				}

				p.expectSemicolon()
				return expr
			}
		} else {
			p.unexpectedToken("Equals")
		}
	} else {
		p.unexpectedToken("Identifier")
	}
	return ast.Nil{}
}

func (p *Parser) parseIf() interface{} {
	var conditionIdentifier = ""
	var ifBlock ast.Expr
	var elseBlock ast.Expr

	// Eat identifier
	if p.eatToken("LParen") {
		if p.eatToken("Identifier") {
			if i, ok := p.Token.TokenType.(lexer.Identifier); ok {
				conditionIdentifier = i.Value
			} else {
				panic("not yet implemented")
			}
		} else {
			p.unexpectedToken("Identifier")
		}

		// Eat right parenthesis for end of the condition
		if p.eatToken("RParen") {
			// Eat left brace for the start of the if block
			if p.eatToken("LBrace") {

				p.advanceToken()
				ifBlock = p.Parse()

				if i, ok := p.TokenStream.Tokens[p.CurrentIndex+1].TokenType.(lexer.Identifier); ok {
					if i.Value == "else" {
						p.advanceToken()

						// Eat left brace for start of the else block
						if p.eatToken("LBrace") {
							p.advanceToken()
							elseBlock = p.Parse()
						} else {
							p.unexpectedToken("LBrace")
						}
					} else {
						elseBlock = ast.Expr{
							Node: nil,
						}
					}
				} else {
					elseBlock = ast.Expr{
						Node: nil,
					}
				}
			} else {
				p.unexpectedToken("LBrace")
			}
		} else {
			p.unexpectedToken("LBrace")
		}
	} else {
		p.unexpectedToken("LParen")
	}
	return ast.If{
		Expr1: ast.Expr{
			Node: ast.Variable{
				Value: conditionIdentifier,
			},
		},
		Expr2: ifBlock,
		Expr3: elseBlock,
	}
}

func (p *Parser) parseCall(identifier string) interface{} {
	var str = ""
	var expr interface{}
	var params []ast.Expr

	// Do While loop for paramaters
	for {
		// Eat String
		if p.eatToken("String") {
			// Create an expression and return it.
			if s, ok := p.Token.TokenType.(lexer.String); ok {
				str = s.Value
			} else {
				panic("not yet implemented")
			}
			params = append(params, ast.Expr{
				Node: ast.Constant{
					Value: ast.String{
						Value: str,
					},
				},
			})
			// Eat identifer
		} else if p.eatToken("Identifier") {
			// Create an expression and return it.
			if i, ok := p.Token.TokenType.(lexer.Identifier); ok {
				str = i.Value
			} else {
				panic("not yet implemented")
			}

			params = append(params, ast.Expr{
				Node: ast.Variable{
					Value: str,
				},
			})
		} else {
			p.unexpectedToken("Identifier or String")
		}

		// Logical check for do while loop
		if !p.eatToken("Comma") {
			break
		}
	}
	expr = ast.Call{
		Value: identifier,
		Exprs: params,
	}
	// Eat RParen
	if p.eatToken("RParen") {
		p.expectSemicolon()
		return expr
	} else {
		p.unexpectedToken("RParen")
		return ast.Nil{}
	}

}

func pop(slice []float64) (float64, []float64) {
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans, slice
}

func (p *Parser) expectSemicolon() {
	if p.eatToken("Semicolon") {
		p.advanceToken()
	} else {
		p.unexpectedToken("Semicolon")
	}
}
