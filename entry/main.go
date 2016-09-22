package main

import (
	"bufio"
	"fmt"
	_ "github.com/MrHuxu/yrel/lexer"
	"github.com/MrHuxu/yrel/parser"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}

func main() {
	// file, err := os.Open("./test.yr")
	// check(err)
	// lexer := lexer.NewLexer(file)
	// fmt.Println(lexer)
	// fmt.Println(lexer.Read())
	// fmt.Println(lexer.Read())
	// fmt.Println(lexer.Read())

	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			parser.YyParse(&parser.Lexer{S: eqn})
		} else {
			break
		}
	}
}
