package interpreter

import (
	"bufio"
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/ast"
	"github.com/kamakuni/go-toy-interpreter/parser"
	"os"
)

type Variable struct {
	Value interface{}
}

type Function struct {
	Value interface{}
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

func (i *Interpreter) Run() {
	node := i.Ast.Node
	i.runBlock(node)
}

func (i *Interpreter) interpretAssign(identifier string, value ast.Expr) {
	if c, ok := value.Node.(ast.Constant); ok {
		i.SymbolTable[identifier] = Symbol{
			SymbolType: Variable,
			Value:      c.Value,
		}
	} else {
		println("Unimplemented feature found!")
	}
}

func (i *Interpreter) interpretCall(identifier string, params []ast.Expr) {
	if identifier == "yaz" {
		i.print(params)
	} else if identifier == "oku" {
		i.get(params)
	}
}

func (i *Interpreter) interpretIf(identifier ast.Expr, ifBlock ast.Expr, elseBlock ast.Expr) {
	var variable = Symbol{
		SymbolType: Variable,
		Value:      ast.String{},
	}

	// Get if condition
	if n, ok := identifier.Node.(Variable); ok {
		i.SymbolTable[n]
	}

	// If condition is a bool value interpret if, otherwise display an error.
	switch v := variable.Value.(type) {
	case ast.Bool:
		// If bool value is true then execute if block.
		if v.Value {
			i.runBlock(ifBlock)
			// If bool value is false and else block is exist, execute else block.
		} else if elseBlock.Node != nil {
			i.runBlock(elseBlock)
		}
	case ast.String:
		panic("Uninitilized variable found!")
	default:
		println("Unimplemented feature found!")
	}
}

func (i *Interpreter) runBlock(expr interface{}) {
	if b, ok := expr.(ast.Block); ok {
		for i, line := range b.Exprs {
			if a, ok := line.Node.(ast.Assign); ok {
				i.interpretAssign(identifier, value)
			} else if c, ok := line.Node.(ast.Call); ok {
				i.interpretCall(identifier, params)
			} else if i, ok := line.Node.(ast.If); ok {
				i.interpretIf(identifier, if_block, else_block)
			} else if e, ok := line.Node.(ast.EOF); ok {
				println("Program has ended.")
			} else {
				println("Unimplemented feature found!")
			}
		}
	} else {
		println("Block not found")
	}
}

func (i *Interpreter) print(params []ast.Expr) {
	var output = ""
	for i, param := range params {
		switch n := param.Node.(type) {
		case ast.Constant:
			switch t := n.Type.(type) {
			case ast.String:
				output += t.Value
			case ast.Number:
				output += fmt.Sprint(t.Value)
			case ast.Bool:
				output += fmt.Sprint(t.Value)
			}
		case ast.Variable:
			if p.SymbolTable[n.Value] != nil {
				output += i.SymbolTable[n.Value].(string)
			} else {
				println("%v variable not found!", n.Value)
			}
		default:
			println("Other type of node found!")
		}
	}
}

func (i *Interpreter) get(params []ast.Expr) {
	for i, param := range params {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		switch v := param.Node.(type) {
		case ast.Variable:
			i.SymbolTable[v.Value] = Symbol{
				SymbolType: Variable,
				Value: ast.String{
					Value: line,
				},
			}
		default:
			println("Parameter requires a variable identifier!")
		}
	}
}
