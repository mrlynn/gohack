package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrlynn/gohack/config"
	"github.com/mrlynn/gohack/storage"
	"github.com/mrlynn/gohack/storage/mongodb"
)

func main() {
	cfg, err := config.GetConfigFromJSON("config.json")

	if err != nil {
		log.Fatal(err)
	}

	client, err := mongodb.NewMongoClient(cfg.Mongo.URI)

	if err != nil {
		log.Fatal(err)
	}

	repository := mongodb.NewMongoRepository(cfg.Mongo.DB, cfg.Mongo.Collection, client)

	storage.SetStorage(repository)

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })
	router.GET("/", IndexPage)
	router.GET("/teams", TeamsPage)
	router.GET("/projects", ProjectsPage)
	router.Run()
}

//IndexPage is a function that displays the main/home page of the hackathon site.
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": config.static.title,
	})
}

//TeamsPage is a function that displays the teams. Each participant will belong to a single team.
func TeamsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": config.static.title,
	})
}

//Projects is a function that displays the main/home page of the hackathon site.
func Projects(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": config.static.title,
	})
}

// Redirect handles POSTs sent to "/"
func Redirect(c *gin.Context) {
	code := c.Param("code")
	url, err := storage.GetURL(code)

	if err != nil {
		log.Println("ERROR:", err)
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"shortUrl": nil,
			"errMsg":   err,
		})
		return
	}
	http.Redirect(c.Writer, c.Request, url, http.StatusMovedPermanently)
}
