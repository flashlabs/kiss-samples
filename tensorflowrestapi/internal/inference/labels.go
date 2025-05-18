package inference

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var Labels []string

func LoadLabels(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	defer func(file *os.File) {
		if e := file.Close(); e != nil {
			log.Println("file.Close", e)
		}
	}(file)

	var labels []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("bufio.Scanner: %w", err)
	}

	Labels = labels

	return nil
}
