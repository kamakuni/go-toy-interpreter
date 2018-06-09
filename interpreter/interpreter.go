package interpreter

import (
	"bufio"
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/ast"
	"github.com/kamakuni/go-toy-interpreter/parser"
	"os"
)

type SymbolType int

const (
	Variable = iota SymbolType
	Function
)

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
		Ast: i.Parse(),
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
				// i.interpretAssign(identifier, value)
			} else if c, ok := line.Node.(ast.Call); ok {
				// TODO
				// i.interpretCall(identifier, params)
			} else if i, ok := line.Node.(ast.If); ok {
				// TODO
				// i.interpretIf(identifier, if_block, else_block)
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

func (i *Interpreter) interpretAssign(identifier string, params ast.Expr) {
	if identifier == "yaz" {
		// TODO
		//i.print(params)
	} else if identifier == "oku" {
		// TODO
		//i.get(params)
	}
}

func (i *Interpreter) print(params []ast.Expr) {
	var output := ""
	for i, param := range params {
		switch n := param.Node.(type){
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
				output += p.SymbolTable[n.Value].(string)
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
		switch v := param.Node.(type){
		case ast.Variable: 
			p.SymbolTable[v.Value] = Symbol{
				SymbolType:Variable,
				Value:ast.String{
					Value: line,
				},
			}
		default: 
			println("Parameter requires a variable identifier!")
		}
	}
}
