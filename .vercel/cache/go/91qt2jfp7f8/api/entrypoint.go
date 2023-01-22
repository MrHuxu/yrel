package main

import (
	"log"
	"net/http"

	"github.com/MrHuxu/yrel/website/templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := templates.GetTemplate("index")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, map[string]any{
		"prd":   true,
		"title": "Yrel",
	})
}
