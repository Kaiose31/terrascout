package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"sync"
)

// Struct to map the image data from JSON
type ImageData struct {
	ID  int    `json:"id"`
	URL string `json:"image_url"`
}

// Function to convert image to RGB format
func convertToRGB(img image.Image) image.Image {
	bounds := img.Bounds()
	rgbImage := image.NewRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			originalColor := img.At(x, y)
			// Get RGBA value
			r, g, b, _ := originalColor.RGBA()
			// Set pixel to RGB values (RGBA() returns values in 0-65535 range, we scale it to 0-255)
			rgbImage.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255})
		}
	}
	return rgbImage
}

// Function to download and save the image as RGB
func downloadAndSaveImage(imageData ImageData) {
	// Get image from URL
	resp, err := http.Get(imageData.URL)
	if err != nil {
		log.Printf("Error downloading image from %s: %v\n", imageData.URL, err)
		return
	}
	defer resp.Body.Close()

	// Decode image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Printf("Error decoding image from %s: %v\n", imageData.URL, err)
		return
	}

	// Convert image to RGB
	rgbImg := convertToRGB(img)

	// Create output file with id.jpg
	outFile, err := os.Create(fmt.Sprintf("../../data/insect_img/%d.jpg", imageData.ID))
	if err != nil {
		log.Printf("Error creating file for image %d: %v\n", imageData.ID, err)
		return
	}
	defer outFile.Close()

	// Save image as JPEG
	err = jpeg.Encode(outFile, rgbImg, nil)
	if err != nil {
		log.Printf("Error encoding image %d: %v\n", imageData.ID, err)
		return
	}

	fmt.Printf("Saved image: %d.jpg\n", imageData.ID)
}

func main() {
	// Open the data.json file
	file, err := os.Open("Insecta.json")
	if err != nil {
		log.Fatalf("Error opening data.json file: %v\n", err)
		return
	}
	defer file.Close()

	// Parse JSON from the file into a slice of ImageData
	var images []ImageData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&images)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v\n", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(images))

	// Loop through the images and download/save each image
	for _, imageData := range images {
		downloadAndSaveImage(imageData)
	}

}
