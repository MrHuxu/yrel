package main

import (
	"fmt"
	"github.com/MrHuxu/yrel"
)

func main() {
	num := &yrel.NumToken{3}
	fmt.Println(yrel.GetTokenType(num))
	fmt.Println(num.GetNumber())
}
