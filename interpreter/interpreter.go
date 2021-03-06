package interpreter

import (
	"bufio"
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/ast"
	"github.com/kamakuni/go-toy-interpreter/parser"
	"os"
)

type Variable struct {
	Value string
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
		Ast:         p.Parse(),
		SymbolTable: map[string]Symbol{},
	}
}

func (i *Interpreter) Run() {
	node := i.Ast.Node
	i.runBlock(node)
}

func (i *Interpreter) interpretAssign(identifier string, value ast.Expr) {
	if c, ok := value.Node.(ast.Constant); ok {
		i.SymbolTable[identifier] = Symbol{
			SymbolType: Variable{},
			Value:      c.Value,
		}
	} else {
		println("Unimplemented feature found!")
	}
}

func (i *Interpreter) interpretCall(identifier string, params []ast.Expr) {
	if identifier == "print" {
		i.print(params)
	} else if identifier == "input" {
		i.get(params)
	}
}

func (i *Interpreter) interpretIf(identifier ast.Expr, ifBlock ast.Expr, elseBlock ast.Expr) {
	var variable Symbol

	// Get if condition
	if n, ok := identifier.Node.(ast.Variable); ok {
		variable = i.SymbolTable[n.Value]
	}

	// If condition is a bool value interpret if, otherwise display an error.
	switch v := variable.Value.(type) {
	case ast.Bool:
		// If bool value is true then execute if block.
		if v.Value {
			if b, ok := ifBlock.Node.(ast.Block); ok {
				i.runBlock(b)
			} else {
				println("Unimplemented feature found!")
			}
			// If bool value is false and else block is exist, execute else block.
		} else if elseBlock.Node != nil {
			if b, ok := elseBlock.Node.(ast.Block); ok {
				i.runBlock(b)
			} else {
				println("Unimplemented feature found!")
			}
		}
	case ast.String:
		panic("Uninitilized variable found!")
	default:
		println("Unimplemented feature found!")
	}
}

func (i *Interpreter) runBlock(expr interface{}) {
	if b, ok := expr.(ast.Block); ok {
		for _, line := range b.Exprs {
			if a, ok := line.Node.(ast.Assign); ok {
				i.interpretAssign(a.Value, a.Expr)
			} else if c, ok := line.Node.(ast.Call); ok {
				i.interpretCall(c.Value, c.Exprs)
			} else if ifNode, ok := line.Node.(ast.If); ok {
				i.interpretIf(ifNode.Expr1, ifNode.Expr2, ifNode.Expr3)
			} else if _, ok := line.Node.(ast.EOF); ok {
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
	for _, param := range params {
		switch n := param.Node.(type) {
		case ast.Constant:
			switch t := n.Value.(type) {
			case ast.String:
				output += t.Value
			case ast.Number:
				output += fmt.Sprint(t.Value)
			case ast.Bool:
				output += fmt.Sprint(t.Value)
			}
		case ast.Variable:
			if s, ok := i.SymbolTable[n.Value].Value.(ast.String); ok {
				output += s.Value
			} else if num, ok := i.SymbolTable[n.Value].Value.(ast.Number); ok {
				output += fmt.Sprint(num.Value)
			} else if b, ok := i.SymbolTable[n.Value].Value.(ast.Bool); ok {
				output += fmt.Sprint(b.Value)
			} else {
				println("%v variable not found!", n.Value)
			}
		default:
			println("Other type of node found!")
		}
	}
	fmt.Println(output)
}

func (i *Interpreter) get(params []ast.Expr) {
	for _, param := range params {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		switch v := param.Node.(type) {
		case ast.Variable:
			i.SymbolTable[v.Value] = Symbol{
				SymbolType: Variable{
					Value: line,
				},
			}
		default:
			println("Parameter requires a variable identifier!")
		}
	}
}
