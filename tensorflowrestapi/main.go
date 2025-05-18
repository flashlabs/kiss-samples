package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashlabs/kiss-samples/tensorflowrestapi/internal/handler"
	"github.com/flashlabs/kiss-samples/tensorflowrestapi/internal/inference"
)

func main() {
	fmt.Println("Loading TF model...")
	if err := inference.LoadModel("model/saved_mobilenet_v2"); err != nil {
		log.Fatalf("Failed to load SavedModel: %v", err)
	}

	fmt.Println("Loading labels...")
	if err := inference.LoadLabels("ImageNetLabels.txt"); err != nil {
		log.Fatalf("Failed to load labels: %v", err)
	}

	fmt.Println("Setting up handlers...")
	http.HandleFunc("/predict", handler.Predict)

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
