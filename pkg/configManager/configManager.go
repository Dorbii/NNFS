package configManager

import (
	"encoding/json"
	"os"
)

type Config struct {
	Inputs  []float32   `json:"inputs"`
	Weights [][]float32 `json:"weights"`
	Biases  []float32   `json:"biases"`
}

var conf *Config

func Read(file string) *Config {
	conf = &Config{}
	f, _ := os.Open(file)
	json.NewDecoder(f).Decode(&conf)
	return conf
}

// func Save(fileName string) {
//If dynamic config is needed
// }
