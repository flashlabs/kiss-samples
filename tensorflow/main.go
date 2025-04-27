package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
	tf "github.com/wamuir/graft/tensorflow"
)

func main() {
	// Load the SavedModel
	model, err := tf.LoadSavedModel("saved_mobilenet_v2", []string{"serve"}, nil)
	if err != nil {
		log.Fatal("LoadSavedModel", err)
	}
	defer func(Session *tf.Session) {
		if e := Session.Close(); e != nil {
			log.Fatal("Session.Close", e)
		}
	}(model.Session)

	// Load an image
	img, err := loadImage("images/1.jpg")
	if err != nil {
		log.Fatal("loadImage", err)
	}

	// Preprocess the image
	tensor, err := makeTensorFromImage(img)
	if err != nil {
		log.Fatal("makeTensorFromImage", err)
	}

	inputOp := model.Graph.Operation("serving_default_x")
	if inputOp == nil {
		log.Fatal("model.Graph.Operation: serving_default_x not found")
	}

	outputOp := model.Graph.Operation("StatefulPartitionedCall")
	if outputOp == nil {
		log.Fatal("model.Graph.Operation: StatefulPartitionedCall not found")
	}

	// Run inference
	outputs, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			inputOp.Output(0): tensor,
		},
		[]tf.Output{
			outputOp.Output(0),
		},
		nil,
	)
	if err != nil {
		log.Fatal("Session.Run", err)
	}

	// Predictions
	predictions := outputs[0].Value().([][]float32)

	// Find the top-1 prediction
	bestIdx := 0
	bestScore := float32(0.0)
	for i, p := range predictions[0] {
		if p > bestScore {
			bestIdx = i
			bestScore = p
		}
	}

	labels, err := loadLabels("ImageNetLabels.txt")
	if err != nil {
		log.Fatal("loadLabels", err)
	}

	fmt.Printf("Predicted label: %s (index: %d, confidence: %.4f)\n", labels[bestIdx], bestIdx, bestScore)
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer func(file *os.File) {
		if e := file.Close(); e != nil {
			log.Fatal("file.Close", e)
		}
	}(file)

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("jpeg.Decode: %w", err)
	}

	return img, nil
}

func makeTensorFromImage(img image.Image) (*tf.Tensor, error) {
	// Resize to 224x224
	resized := resize.Resize(224, 224, img, resize.Bilinear)

	// Create a 4D array to hold input
	bounds := resized.Bounds()
	batch := make([][][][]float32, 1) // batch size 1
	batch[0] = make([][][]float32, bounds.Dy())

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		row := make([][]float32, bounds.Dx())
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := resized.At(x, y).RGBA()
			row[x] = []float32{
				float32(r) / 65535.0, // normalize to [0,1]
				float32(g) / 65535.0,
				float32(b) / 65535.0,
			}
		}
		batch[0][y] = row
	}

	return tf.NewTensor(batch)
}

func loadLabels(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer func(file *os.File) {
		if e := file.Close(); e != nil {
			log.Fatal("file.Close", e)
		}
	}(file)

	var labels []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("bufio.Scanner: %w", err)
	}

	return labels, nil
}
