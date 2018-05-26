package ast

import (
	"github.com/kamakuni/go-toy-interpreter/lexer"
)

type Expr struct {
	Span lexer.Span
	Node Expr_
}

type Expr_ struct {
	Type  ExprType
	Value interface{}
}

type ExprType int

const (
	// Block of statements
	Block ExprType = iota
	// Add two expressions.
	Add
	// Subtract two expressions
	Sub
	// Multiply two expressions
	Mul
	// Divide two expressions
	Div
	// Variable expression
	Variable
	// Constant expression
	Constant
	// Assignment expression
	Assign
	// If expression 'if expr { expr } else { expr }'
	If
	// Function Call, first field is name of the function, second is list of arguments
	Call
	// Literal expression
	Literal
	// End of File
	EOF
	// Null
	Nil
)

type ConstantType int

const (
	String ConstantType = iota
	Number
	Bool
)
