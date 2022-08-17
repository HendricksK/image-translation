package main

import (
	"fmt"
	"net/http"
	"os"

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
	router.GET("/images/:topic", getImagesBasedOnTopic)
	return router
}

func main() {
	router := setupRouter()
	port := os.Getenv("PORT")

	if port != "" {
		router.Run("localhost:" + port)
	} else {
		router.Run("localhost:5000")
	}
}

func getImageTopics(c *gin.Context) {
	data := imageimporter.GetAllUploadedImageTopics()
	c.IndentedJSON(http.StatusOK, data)
}

// https://github.com/gin-gonic/gin/issues/2555
func getImagesBasedOnTopic(c *gin.Context) {

	var imageReturn image

	topic := c.Param("topic")
	fmt.Println(topic)
	// Get topic from gin context, if no topic
	// bad request
	if topic == "" {
		c.IndentedJSON(http.StatusBadRequest, "")
	}

	data := imageimporter.GetImagesBasedOnTopic(topic)

	for _, img := range data {
		if img != "" {
			imageReturn.Images = append(imageReturn.Images, img)
		}
	}

	imageReturn.RequestedTopic = topic
	imageReturn.Number = len(imageReturn.Images)

	c.IndentedJSON(http.StatusOK, imageReturn)
}
