package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
	"os"
	"time"
)

type Post struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Content   string
	CreatedAt time.Time
	Appends   *[]Append
	Replies   *[]Reply
}

type Append struct {
	PostID    bson.ObjectId
	Text      string
	CreatedAt time.Time
}

type Reply struct {
	PostID    bson.ObjectId
	Text      string
	ReplyTo   int
	CreatedAt time.Time
}

func getAllPosts(c *gin.Context, db *mgo.Database) {
	postCollection := db.C("post")
	appendCollection := db.C("append")
	replyCollection := db.C("reply")

	var posts []Post
	err := postCollection.Find(bson.M{}).All(&posts)
	if err != nil {
		fmt.Println(err)
	}

	for index := range posts {
		var appends []Append
		err = appendCollection.Find(bson.M{"postid": posts[index].ID}).All(&appends)
		if err != nil {
			fmt.Println(err)
		}
		posts[index].Appends = &appends

		var replies []Reply
		err = replyCollection.Find(bson.M{"postid": posts[index].ID}).All(&replies)
		if err != nil {
			fmt.Println(err)
		}
		posts[index].Replies = &replies
	}

	c.JSON(200, gin.H{
		"result": "success",
		"posts":  &posts,
	})
}

func createPost(c *gin.Context, db *mgo.Database) {
	var newPost Post
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&newPost)
	if err != nil {
		fmt.Println(err)
	}
	newPost.ID = bson.NewObjectId()
	newPost.CreatedAt = time.Now()

	post := db.C("post")
	err = post.Insert(&newPost)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"result":  "success",
		"newPost": &newPost,
	})
}

func appendPost(c *gin.Context, db *mgo.Database) {
	var newAppend Append
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&newAppend)
	if err != nil {
		fmt.Println(err)
	}
	newAppend.CreatedAt = time.Now()

	append := db.C("append")
	err = append.Insert(&newAppend)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"result":    "success",
		"newAppend": &newAppend,
	})
}

func replyPost(c *gin.Context, db *mgo.Database) {
	var newReply Reply
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&newReply)
	if err != nil {
		fmt.Println(err)
	}
	newReply.CreatedAt = time.Now()

	reply := db.C("reply")
	err = reply.Insert(&newReply)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"result":   "success",
		"newReply": &newReply,
	})
}

func main() {
	fmt.Println("==> 🌎  Listening on port 8082. Open up http://localhost:8082/ in your browser.")

	// setup database
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	barDB := session.DB("bar")

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

	postRoutes := router.Group("/post")
	{
		postRoutes.GET("/", func(c *gin.Context) { getAllPosts(c, barDB) })
		postRoutes.POST("/create", func(c *gin.Context) { createPost(c, barDB) })
		postRoutes.POST("/append", func(c *gin.Context) { appendPost(c, barDB) })
		postRoutes.POST("/reply", func(c *gin.Context) { replyPost(c, barDB) })
	}

	router.Run(":8082")
}
