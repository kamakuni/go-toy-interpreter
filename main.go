package main

import (
	"fmt"
	"github.com/kamakuni/go-toy-interpreter/lexer"
	"os"
)

func main() {
	var path = "src/test/main.c"
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Your source file is: ", os.Args[1])
		path = os.Args[1]
	} else {
		fmt.Println("You are using default test source path: src/test/main.c")
	}
	ts := lexer.CreateTokenStream("this is codes")
	fmt.Println(ts)
}
