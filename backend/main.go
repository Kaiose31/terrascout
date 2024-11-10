package main

import (
	"backend/wildlife"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var db []wildlife.Observation

func createDB() {
	content, err := os.ReadFile("../data/observations.json")
	if err != nil {
		log.Fatal("Failed to read observation data: ", err)
	}

	// err = json.Unmarshal(content, &db)
	// if err != nil {
	// 	log.Fatal("Error parsing json: ", err)
	// }

}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")

	})

	r.GET("/refresh", func(ctx *gin.Context) {
		db = nil
		createDB()
		ctx.String(http.StatusOK, "refreshed db")
	})

	r.POST("/map", func(c *gin.Context) {

		//Create bounding box
		strings.Split(c.Query("box"), ",")
		//filter data for given bounding box
	})

	return r
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
