package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	ts := CreateTokenStream("this is codes")
	fmt.Println(ts)
}
