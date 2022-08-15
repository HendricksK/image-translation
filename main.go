package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type image struct {
	RequestedTopic string   `json:"requestedTopic"`
	Images         []string `json:"Images"`
	Number         int      `json:"Number"`
}

var images = []image{
	{
		RequestedTopic: "the+room+memes",
		Images: []string{
			"https://preview.redd.it/0ndypduzlf641.jpg?width=640&crop=smart&auto=webp&s=5cf49a2deb07cfa5bf628e1dfb074b0452bb8041",
			"https://preview.redd.it/q6ez8ub45i931.jpg?width=640&crop=smart&auto=webp&s=4b4859c1cc752ac3310496852b8fd19a77aa7fc7",
		},
		Number: 2,
	},
}

var topics = []string{
	"the+room+memes",
	"random",
	"all",
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/topics", getImageTopics)
	router.GET("/images", getImagesBasedOnTopic)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}

func getImageTopics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, topics)
}

// mock data for now
func getImagesBasedOnTopic(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, images)
}
