package configx

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func NewConfig[T any](path string) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	config, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var cfg T
	err = yaml.Unmarshal(config, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
