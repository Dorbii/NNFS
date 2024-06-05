package configManager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type JConfig struct {
	Inputs  []float32 `json:"inputs"`
	Weights []float32 `json:"weights"`
	Biases  []float32 `json:"biases"`
}

type YConfig struct {
	Section  string    `yaml:"section"`
	Inputs   []float32 `yaml:"inputs"`
	Weights  []float32 `yaml:"weights"`
	Biases   []float32 `yaml:"biases"`
	Weights2 []float32 `yaml:"weights2,omitempty"`
	Biases2  []float32 `yaml:"biases2,omitempty"`
}

func ReadJSON(file string) *JConfig {
	c := &JConfig{}
	f, err := os.Open(file)
	if err != nil {
		msg := fmt.Errorf("error opening json file: %w", err)
		log.Error(msg)
		return c
	}
	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		msg := fmt.Errorf("error decoding json file: %w", err)
		log.Error(msg)
		return nil
	}
	return c
}

func ReadYAML(file string) map[string]YConfig {
	m := make(map[string]YConfig)
	f, err := os.ReadFile(file)
	if err != nil {
		msg := fmt.Errorf("error reading yaml file: %w", err)
		log.Fatal(msg)
	}
	decoder := yaml.NewDecoder(bytes.NewBuffer(f))

	for {
		c := YConfig{}
		if err := decoder.Decode(&c); err != nil {
			if err == io.EOF {
				msg := fmt.Errorf("hit EOF flag: %w", err)
				log.Info(msg)
				break
			}
			msg := fmt.Errorf("failed to decode: %w", err)
			log.Error(msg)
		}
		m[c.Section] = c
	}
	return m
}
