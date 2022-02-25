package mockspec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// MockSpecification specification of mock
type MockSpecification struct {
	Body       *string                `json:"body"`
	JSONBody   map[string]interface{} `json:"jsonBody"`
	Base64Body []byte                 `json:"base64Body"`
	Path       string                 `json:"path"`
	Status     int                    `json:"status"`
	Headers    http.Header            `json:"headers"`
}

// CollectFromDirectory collect specifications from the specified directory
func CollectFromDirectory(dir string) ([]MockSpecification, error) {
	var specs []MockSpecification
	if err := filepath.Walk(dir, createDirWalker(&specs)); err != nil {
		return nil, errors.Wrap(err, "error reading specification")
	}
	return specs, nil
}

func createDirWalker(specCollector *[]MockSpecification) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("[%s] error opening file '%w'", path, err)
			}
			read, err := ioutil.ReadAll(file)
			if err != nil {
				return fmt.Errorf("[%s] error reading file'%w'", path, err)
			}

			var spec MockSpecification
			err = json.Unmarshal(read, &spec)
			if err != nil {
				return fmt.Errorf("[%s] error unmarshalling specification '%w'", path, err)
			}
			if (spec.JSONBody != nil && spec.Body != nil) ||
				(spec.Body != nil && spec.Base64Body != nil) ||
				(spec.JSONBody != nil && spec.Base64Body != nil) {
				return fmt.Errorf("[%s] only one of the jsonBody, body, base64Body can be specified", path)
			}
			if len(spec.Path) == 0 {
				return fmt.Errorf("[%s] path is required argument", path)
			}
			*specCollector = append(*specCollector, spec)
			return nil
		}
		return nil
	}
}
