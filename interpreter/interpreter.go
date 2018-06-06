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

func (i *Interpreter) runBlock(expr interface{}) {
	if b, ok := expr.(ast.Block); ok {
		for i, line := range b.Exprs {
			if a, ok := line.Node.(ast.Assign); ok {
				// TODO
				// p.interpretAssign(identifier, value)
			} else if c, ok := line.Node.(ast.Call); ok {
				// TODO
				// p.interpretCall(identifier, params)
			} else if i, ok := line.Node.(ast.If); ok {
				// TODO
				// p.interpretIf(identifier, if_block, else_block)
			} else if e, ok := line.Node.(ast.EOF); ok {
				// TODO
				println("Program has ended.")
			} else {
				println("Unimplemented feature found!")
			}
		}
	} else {
		println("")
	}
}
