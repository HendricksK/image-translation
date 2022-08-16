package main

import (
	"net/http"

	"github.com/HendricksK/image-translation/imageimporter"
	"github.com/gin-gonic/gin"
)

type image struct {
	RequestedTopic string   `json:"requestedTopic"`
	Images         []string `json:"Images"`
	Number         int      `json:"Number"`
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
	data := imageimporter.GetAllUploadedImageTopics()
	c.IndentedJSON(http.StatusOK, data)
}

func getImagesBasedOnTopic(c *gin.Context) {

	var imageReturn image
	data := imageimporter.GetImagesBasedOnTopic("the+room+memes")

	for _, img := range data {
		if img != "" {
			imageReturn.Images = append(imageReturn.Images, img)
		}
	}

	imageReturn.RequestedTopic = "the+room+memes"
	imageReturn.Number = len(imageReturn.Images)

	c.IndentedJSON(http.StatusOK, imageReturn)
}
