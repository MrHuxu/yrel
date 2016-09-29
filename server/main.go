package main

import (
	"fmt"
	"github.com/MrHuxu/yrel/parser"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func submitCode(c *gin.Context) {
	str, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		lexer := parser.Lexer{Input: string(str)}
		parser.YyParse(&lexer)
		for _, stat := range parser.Statements {
			stat.Execute()
		}
		fmt.Println(parser.Tokens, parser.Statements, parser.Outputs)
	}
}

func main() {
	fmt.Println("==> 🌎  Listening on port 8082. Open up http://localhost:8082/ in your browser.")

	// init router
	var router *gin.Engine
	if "Production" == os.Getenv("ENV") {
		logFile, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("error opening log file")
		}
		defer logFile.Close()

		gin.DefaultWriter = io.Writer(logFile)
		gin.SetMode(gin.ReleaseMode)

		router = gin.New()
		router.Use(gin.Logger())
		router.StaticFile("./bundle.js", "./built/bundle.js")
	} else {
		gin.SetMode(gin.DebugMode)
		router = gin.New()
	}
	router.LoadHTMLGlob("templates/*")

	// setup routers
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"prd":   "Production" == os.Getenv("ENV"),
			"title": "Yrel",
		})
	})

	routes := router.Group("/yrel")
	{
		routes.POST("/", submitCode)
	}

	router.Run(":8082")
}
