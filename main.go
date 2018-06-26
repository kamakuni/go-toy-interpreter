package main

import (
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/interpreter"
	"github.com/kamakuni/go-toy-interpreter/lexer"
	"github.com/kamakuni/go-toy-interpreter/parser"
	"io/ioutil"
	"os"
)

func main() {
	// Try to get file from arguments.
	// If given run this file, otherwise run "src/test/main.c"
	var path = "test.txt"
	fmt.Println(os.Args)

	// For custom source file for interpreting.
	if len(os.Args) > 1 {
		fmt.Println("Your source file is: ", os.Args[1])
		path = os.Args[1]
	} else {
		fmt.Println("You are using default test source path: src/test/main.c")
	}
	// Open the source file.
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("couldn't open file")
	}
	// Get Tokens from the Lexer Module
	ts := lexer.CreateTokenStream(string(data))
	// Creating a new Parser instance for AST.
	parser := parser.CreateParser(ts, lexer.Span{})
	// Creating an interpreter module for program.
	program := interpreter.NewInterpreter(parser)
	program.Run()
}
