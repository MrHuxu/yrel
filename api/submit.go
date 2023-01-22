package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MrHuxu/yrel/parser"
)

// Submit ...
func Submit(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	lexer := parser.Lexer{Input: string(bytes)}
	parser.YyParse(&lexer)
	for _, stat := range parser.Statements {
		stat.Execute()
	}

	json.NewEncoder(w).Encode(map[string]any{
		"result": "success",
		"content": map[string]interface{}{
			"tokens":     parser.Tokens,
			"statements": parser.Statements,
			"outputs":    parser.Outputs,
		},
	})
}
