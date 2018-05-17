package main

import (
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/ast"
	"os"
)

func main() {
	fmt.Println(os.Args)
	ts := ast.CreateTokenStream("this is codes")
	fmt.Println(ts)
}
