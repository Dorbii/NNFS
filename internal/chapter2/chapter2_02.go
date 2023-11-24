package chapter2

import (
	"Dorbii/NNFS/pkg/configManager"
	"fmt"

	t "gorgonia.org/tensor"
)

type config struct {
	Inputs  []float32   `json:"inputs"`
	Weights [][]float32 `json:"weights"`
	Biases  []float32   `json:"biases"`
}

func chapter2_02() {
	configManager.Read("../../configs/config.json", &config{})
	inputs := t.New(t.WithBacking(&config.Inputs))
	fmt.Println(inputs)
	//layerOutputs := t.Dot(t.New(inputs), t.New(weights))
}
