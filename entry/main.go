package main

import (
	_ "bufio"
	"fmt"
	_ "github.com/MrHuxu/yrel/lexer"
	"github.com/MrHuxu/yrel/parser"
	_ "os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func readline(fi *bufio.Reader) (string, bool) {
// 	s, err := fi.ReadString('\n')
// 	if err != nil {
// 		return "", false
// 	}
// 	return s, true
// }

func main() {
	fmt.Println("Yrel 0.0.1  Copyright (C) 2016-2018 xhu.me, Xu Hu")
	var input = "a = 1\n" +
		"print a"
	parser.YyParse(&parser.Lexer{S: input})
	for _, stat := range parser.Statements {
		stat.Execute()
	}
	// fi := bufio.NewReader(os.NewFile(0, "stdin"))

	// for {
	// 	var eqn string
	// 	var ok bool

	// 	fmt.Printf("> ")
	// 	if eqn, ok = readline(fi); ok {
	// 		parser.YyParse(&parser.Lexer{S: eqn})
	// 	} else {
	// 		break
	// 	}
	// }
}
