package main

import (
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/lexer"
	"os"
)

func main() {
	fmt.Println(os.Args)
	ts := lexer.CreateTokenStream("this is codes")
	fmt.Println(ts)
}
