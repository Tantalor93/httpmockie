package mockspec

import (
	"encoding/json"
	"go-http-mock/internal/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MockSpecification struct {
	JsonBody map[string]interface{} `json:"jsonBody"`
	Body     *string                 `json:"body"`
	Path     string                 `json:"path"`
	Status   int                    `json:"status"`
}

func CollectFromDirectory(dir string) []MockSpecification {
	var specs []MockSpecification
	if err := filepath.Walk(dir, createDirWalker(&specs)); err != nil {
		panic(err)
	}
	return specs
}

func createDirWalker(specCollector *[]MockSpecification) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			log.Logger.Infof("Reading mock specification '%s'", path)

			file, e := os.Open(path)
			if e != nil {
				log.Logger.Errorf("Error opening while '%s'", path)
				return nil
			}
			read, err := ioutil.ReadAll(file)
			if err != nil {
				log.Logger.Errorf("Error while reading file '%s'", path)
				return nil
			}

			var spec MockSpecification
			e = json.Unmarshal(read, &spec)
			if e != nil {
				log.Logger.Errorf("Error while parsing mock specification in file '%s'", path)
				return nil
			}
			if spec.JsonBody != nil && spec.Body != nil {
				log.Logger.Errorf("jsonBody and body cannot be used at the same time '%s'", path)
				return nil
			}
			*specCollector = append(*specCollector, spec)
			return nil
		}
		return nil
	}
}
