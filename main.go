// main.go
package main

import (
	"favourite-artist-service/app/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/get-info", func(c *gin.Context) {
		region := c.Query("region")
		trackName := c.Query("track_name")
		artistName := c.Query("artist_name")

		result, err := handlers.GetInformation(region, trackName, artistName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
