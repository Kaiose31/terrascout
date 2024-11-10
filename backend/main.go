package main

import (
	"backend/utils"
	"backend/wildlife"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var db []wildlife.Observation

func createDB() {
	content, err := os.ReadFile("../data/obs.json")
	if err != nil {
		log.Fatal("Failed to read observation data: ", err)
	}

	err = json.Unmarshal(content, &db)
	if err != nil {
		log.Fatal("Error parsing json: ", err)
	}

}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	//clear and reset db
	r.GET("/refresh", func(ctx *gin.Context) {
		db = nil
		createDB()
		ctx.String(http.StatusOK, "refreshed db")
	})

	r.GET("/map", func(c *gin.Context) {

		//Create bounding box
		coords, err := utils.ConvertStringsToFloats(strings.Split(c.Query("box"), ","))
		if err != nil {
			log.Fatal("Failed to parse box")
		}
		//filter db for given bounding box
		utils.FilterObservations(&db, coords[0], coords[1], coords[2], coords[3])

		//fetch summary for taxon
		c.String(http.StatusOK, "Filtered data according to bounding box")
	})

	//forward given image and possibly
	r.POST("/match", func(ctx *gin.Context) {

		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		imageFile, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
			return
		}

		defer imageFile.Close()

		imageData := make([]byte, file.Size)
		_, err = imageFile.Read(imageData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the file"})
			return
		}

		result, err := utils.RunML(imageData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ML processing failed: %v", err)})
			return
		}

		//send the data for the recognized species
		ctx.JSON(http.StatusOK, gin.H{"result": result})

	})

	return r
}

func main() {
	go createDB()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
