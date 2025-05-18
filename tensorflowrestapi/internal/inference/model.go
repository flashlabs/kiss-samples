package inference

import (
	"fmt"

	tf "github.com/wamuir/graft/tensorflow"
)

var Model *tf.SavedModel

func LoadModel(path string) (err error) {
	Model, err = tf.LoadSavedModel(path, []string{"serve"}, nil)
	if err != nil {
		return fmt.Errorf("LoadSavedModel: %w", err)
	}

	return nil
}
