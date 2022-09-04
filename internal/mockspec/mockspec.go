package mockspec

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

//go:embed specification.json
var jsonSchema string

// MockSpecification specification of mock.
type MockSpecification struct {
	Body       *string                `json:"body"`
	JSONBody   map[string]interface{} `json:"jsonBody"`
	Base64Body []byte                 `json:"base64Body"`
	Path       string                 `json:"path"`
	Status     int                    `json:"status"`
	Headers    http.Header            `json:"headers"`
	Delay      *Delay                 `json:"delay"`
}

// Delay specifies delay of single endpoint.
type Delay struct {
	DurationMs  int `json:"durationMs"`
	DeviationMs int `json:"deviationMs"`
}

// CollectFromDirectory collect specifications from the specified directory.
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
			read, err := io.ReadAll(file)
			if err != nil {
				return fmt.Errorf("[%s] error reading file'%w'", path, err)
			}

			schemaLoader := gojsonschema.NewStringLoader(jsonSchema)
			documentLoader := gojsonschema.NewStringLoader(string(read))
			res, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil {
				return fmt.Errorf("[%s] error while validating schema, '%w'", path, err)
			}
			if !res.Valid() {
				for _, v := range res.Errors() {
					return fmt.Errorf("[%s] %s", path, v.String())
				}
			}

			var spec MockSpecification
			err = json.Unmarshal(read, &spec)
			if err != nil {
				return fmt.Errorf("[%s] error while unmarshalling JSON '%w'", path, err)
			}

			*specCollector = append(*specCollector, spec)
			return nil
		}
		return nil
	}
}
