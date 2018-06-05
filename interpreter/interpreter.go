package interpreter

import (
	"github.com/kamakuni/go-toy-interpreter/ast"
	"github.com/kamakuni/go-toy-interpreter/parser"
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

func NewInterpreter(p parser.Parser) Interpreter {
	return Interpreter{
		Ast: p.Parse(),
	}
}

func (i *Interpreter) run() {
	node := i.Ast.Node
	i.runBlock(node)
}
