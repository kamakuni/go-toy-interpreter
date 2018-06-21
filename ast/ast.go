package ast

import (
	"github.com/kamakuni/go-toy-interpreter/lexer"
)

type Expr struct {
	Span lexer.Span
	//	Node Expr_
	Node interface{}
}

/*
type Expr_ struct {
	Type  ExprType
	Value interface{}
}
*/
type Expr_ int

/*const (
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
)*/

type Block struct {
	Expr_
	Exprs []Expr
}

// Add two expressions.
type Add struct {
	Expr_
	Expr1 Expr
	Expr2 Expr
}

// Subtract two expressions
type Sub struct {
	Expr_
	Expr1 Expr
	Expr2 Expr
}

// Multiply two expressions
type Mul struct {
	Expr_
	Expr1 Expr
	Expr2 Expr
}

// Divide two expressions
type Div struct {
	Expr_
	Expr1 Expr
	Expr2 Expr
}

// Variable expression
type Variable struct {
	Expr_
	Value string
}

// Constant expression
type Constant struct {
	Expr_
	Value interface{}
}

// Assignment expression
type Assign struct {
	Expr_
	Value string
	Expr  Expr
}

// If expression 'if expr { expr } else { expr }'
type If struct {
	Expr_
	Expr1 Expr
	Expr2 Expr
	Expr3 Expr
}

// Function Call, first field is name of the function, second is list of arguments
type Call struct {
	Expr_
	Value string
	Exprs []Expr
}

// Literal expression
type Literal struct {
	Expr_
	Value float64
}

// End of File
type EOF struct {
	Expr_
}

// Null
type Nil struct {
	Expr_
}

type ConstantType int

type String struct {
	ConstantType
	Value string
}

type Number struct {
	ConstantType
	Value float64
}

type Bool struct {
	ConstantType
	Value bool
}
