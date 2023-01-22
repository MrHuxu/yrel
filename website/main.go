package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/MrHuxu/yrel/parser"

	"github.com/gin-gonic/gin"
)

func submitCode(c *gin.Context) {
	parser.Tokens = parser.Tokens[:0]
	parser.Statements = parser.Statements[:0]
	parser.Outputs = parser.Outputs[:0]

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"result":  "failure",
			"content": map[string]interface{}{},
		})
	} else {
		lexer := parser.Lexer{Input: string(bytes)}
		parser.YyParse(&lexer)
		for _, stat := range parser.Statements {
			stat.Execute()
		}
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
			"content": map[string]interface{}{
				"tokens":     parser.Tokens,
				"statements": parser.Statements,
				"outputs":    parser.Outputs,
			},
		})
	}
}

func main() {
	fmt.Println("==> 🌎  Listening on port 8082. Open up http://localhost:8082/ in your browser.")

	// init router
	var router *gin.Engine
	if "release" == os.Getenv("GIN_MODE") {
		logFile, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("error opening log file")
		}
		defer logFile.Close()

		gin.DefaultWriter = io.Writer(logFile)

		router = gin.New()
		router.Use(gin.Logger())
		router.StaticFile("./built/bundle.js", "./built/bundle.js")
	} else {
		gin.SetMode(gin.DebugMode)
		router = gin.New()
	}
	router.LoadHTMLGlob("templates/*")

	// setup routers
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"prd":   "release" == os.Getenv("GIN_MODE"),
			"title": "Yrel",
		})
	})
	router.POST("/submit", submitCode)

	router.Run(":8082")
}
