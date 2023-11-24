package configManager

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Inputs  []float32 `json:"inputs"`
	Weights []float32 `json:"weights"`
	Biases  []float32 `json:"biases"`
}

var conf *Config

func Read(file string) *Config {
	conf = &Config{}
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return conf
	}
	json.NewDecoder(f).Decode(&conf)
	return conf
}

// func Save(fileName string) {
//If dynamic config is needed
// }
