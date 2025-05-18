package handler

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/nfnt/resize"
	tf "github.com/wamuir/graft/tensorflow"

	"github.com/flashlabs/kiss-samples/tensorflowrestapi/internal/inference"
)

func Predict(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get images", http.StatusBadRequest)

		return
	}
	defer func(file multipart.File) {
		if e := file.Close(); e != nil {
			log.Println("file.Close", e)
		}
	}(file)

	img, err := jpeg.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusBadRequest)

		return
	}

	tensor, err := makeTensorFromImage(img)
	if err != nil {
		http.Error(w, "Failed to make tensor from image", http.StatusInternalServerError)

		return
	}

	input := inference.Model.Graph.Operation("serving_default_x")
	output := inference.Model.Graph.Operation("StatefulPartitionedCall")

	outputs, err := inference.Model.Session.Run(
		map[tf.Output]*tf.Tensor{
			input.Output(0): tensor,
		},
		[]tf.Output{
			output.Output(0),
		},
		nil,
	)
	if err != nil {
		http.Error(w, "Failed to run inference", http.StatusInternalServerError)

		return
	}

	predictions := outputs[0].Value().([][]float32)

	bestIdx, bestScore := 0, float32(0.0)
	for i, p := range predictions[0] {
		if p > bestScore {
			bestIdx, bestScore = i, p
		}
	}

	label := inference.Labels[bestIdx]

	_, err = fmt.Fprintf(w, `{"class_id": %d, "label": "%s", "confidence": %.4f}`+"\n", bestIdx, label, bestScore)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)

		return
	}
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
