package main

import (
	_ "fmt"
	"github.com/MrHuxu/yrel"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./test.yr")
	check(err)
	lexer := yrel.NewLexer(file)
	lexer.Read()
}
