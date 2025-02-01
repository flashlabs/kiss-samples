package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/xeipuuv/gojsonschema"
)

var ErrInvalidPayload = errors.New("invalid payload")

type RequestPayload struct {
	ID uuid.UUID `json:"id"`
}

func main() {
	// Request payload.
	p, err := payload()
	if err != nil {
		log.Fatal(err)
	}

	// Validate incoming payload.
	if err = ValidateSchema(p, "schema/payload.json"); err != nil {
		log.Fatal(err)
	}

	// We're good!
	fmt.Println("Schema validated")
}

// ValidateSchema validates payload against JSON schema located in the schema file.
func ValidateSchema(payload map[string]any, schema string) error {
	// Let's read the schema file
	file, err := os.ReadFile(schema)
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}

	// Prepare validators for schema...
	schemaLoader := gojsonschema.NewStringLoader(string(file))
	// ... and payload.
	payloadLoader := gojsonschema.NewGoLoader(payload)

	// Validate schema against payload:
	result, err := gojsonschema.Validate(schemaLoader, payloadLoader)
	if err != nil {
		return fmt.Errorf("gojsonschema.Validate: %w", err)
	}

	// If there was something wrong, communicate the errors:
	if !result.Valid() {
		errMsg := "JSON validation failed:\n"
		for _, desc := range result.Errors() {
			errMsg += fmt.Sprintf("- %s\n", desc)
		}
		return fmt.Errorf("%w: %s", ErrInvalidPayload, errMsg)
	}

	return nil
}

// payload returns sample request payload in map[string]any format.
func payload() (map[string]any, error) {
	r := &RequestPayload{ID: uuid.New()}

	payload, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	var data map[string]any
	if err := json.Unmarshal(payload, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return data, nil
}
