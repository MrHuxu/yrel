package main

import (
	"fmt"
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
	fmt.Println(lexer)
	fmt.Println(lexer.Read())
	fmt.Println(lexer.Read())
	fmt.Println(lexer.Read())
}
