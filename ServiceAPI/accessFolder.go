package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func imageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Path[len("/image/"):]       // Extract image name from URL path
	imagePath := ROOT_PATH + "/image/" + imageName // Build image path

	// Read the image file
	imgBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Fprintf(w, "Error reading image: %v", err)
		return
	}

	// Get image content type based on file extension
	contentType := http.DetectContentType(imgBytes)

	// Set content type header
	w.Header().Set("Content-Type", contentType)

	// Write image data to response body
	w.Write(imgBytes)
}
