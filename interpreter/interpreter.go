package interpreter

import (
	"github.com/kamakuni/go-toy-interpreter/ast"
)

type Variable struct {
}

type Function struct {
}

type Symbol struct {
	SymbolType interface{}
	Value      interface{}
}

type Interpreter struct {
	Ast         ast.Expr
	SymbolTable map[string]Symbol
}
