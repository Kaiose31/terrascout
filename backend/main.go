package main

import (
	"backend/utils"
	"backend/wildlife"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db []wildlife.Observation
var taxa []wildlife.Taxon

func createDB() {
	content, err := os.ReadFile("../data/obs_large.json")
	if err != nil {
		log.Fatal("Failed to read observation data: ", err)
	}

	err = json.Unmarshal(content, &db)
	if err != nil {
		log.Fatal("Error parsing json: ", err)
	}

	taxon, err := os.ReadFile("../data/taxon_large.json")
	if err != nil {
		log.Fatal("Failed to read observation data: ", err)
	}

	err = json.Unmarshal(taxon, &taxa)
	if err != nil {
		log.Fatal("Error parsing json: ", err)
	}

}

type ImageRequest struct {
	Image string `json:"image"`
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(cors.Default()) // This allows all origins by default

	//clear and reset db
	r.GET("/refresh", func(ctx *gin.Context) {
		db = nil
		createDB()
		ctx.JSON(http.StatusOK, gin.H{"records": len(db), "taxa": len(taxa)})
	})

	r.GET("/obs", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": db[:100]})
	})

	r.POST("/match", func(ctx *gin.Context) {

		var req ImageRequest

		// Bind the JSON body to the ImageRequest struct
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		// Check if the "image" key is present
		if req.Image == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No image data found"})
			return
		}

		// Decode the base64 image string
		// Remove the data URL scheme prefix (if it exists) to get the base64 string
		base64Image := req.Image
		base64Image = strings.TrimPrefix(base64Image, "data:image/jpeg;base64,")

		// Decode the base64 string into bytes
		imageData, err := base64.StdEncoding.DecodeString(base64Image)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64 image"})
			return
		}

		// imageFile, err := file.Open()
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
		// 	return
		// }

		// defer imageFile.Close()

		// imageData := make([]byte, file.Size)
		// _, err = imageFile.Read(imageData)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the file"})
		// 	return
		// }

		scientificName, err := utils.RunML(imageData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ML processing failed: %v", err)})
			return
		}
		fmt.Println(scientificName)
		//find the Scientific name in the db
		utils.FilterObsByName(&db, scientificName)

		if len(db) == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not Identify Species"})
			return
		}
		fmt.Println(db[0].TaxonID)
		//fetch details

		utils.FilterTaxon(&taxa, db[0].TaxonID)

		if len(taxa) == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch Wiki"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"obs": db[0], "wiki": taxa[0].WikipediaSummary})

	})

	return r
}

func main() {
	createDB()
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
