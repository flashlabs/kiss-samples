package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	filePath   = "../example.jpg"
	yoloAPIURL = "http://localhost:8000/detect"
)

// main is the entry point for the application. It prepares the image,
// sends it to the YOLO API, and prints the result.
func main() {
	// Prepare the image file as a multipart form
	body, contentType, err := prepareMultipartForm(filePath)
	if err != nil {
		log.Fatal("Error preparing multipart form: ", err)
	}

	// Send the HTTP POST request to the YOLO API
	respBytes, err := sendYOLORequest(yoloAPIURL, body, contentType)
	if err != nil {
		log.Fatal("Error sending YOLO request: ", err)
	}

	// Print the detection results
	fmt.Println(string(respBytes))
}

// prepareMultipartForm creates a multipart/form-data body from the given file path.
// It returns the form body, content type, and any error encountered.
func prepareMultipartForm(filePath string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if e := file.Close(); e != nil {
			log.Println("Failed to close file", e)
		}
	}()

	// Create a new form file field
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, "", fmt.Errorf("failed to create form file: %w", err)
	}

	// Copy the image data into the form
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, "", fmt.Errorf("failed to copy file: %w", err)
	}

	// Close the multipart writer
	if err = writer.Close(); err != nil {
		log.Println("Failed to close writer", err)
	}

	return body, writer.FormDataContentType(), nil
}

// sendYOLORequest sends the image as a multipart POST request to the specified YOLO API.
// It returns the response body or an error.
func sendYOLORequest(apiURL string, body *bytes.Buffer, contentType string) ([]byte, error) {
	// Create a new HTTP POST request with the multipart data
	req, err := http.NewRequest(http.MethodPost, apiURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", contentType)

	// Send the request and get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			log.Println("Failed to close body", e)
		}
	}()

	// Read and return the response body
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBytes, nil
}
