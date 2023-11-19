package chapter2

import (
	"fmt"

	"github.com/Dorbii/NNFS/pkg/goZip"
)

func Chapter2() string {
	inputs := []float32{1.0, 2.0, 3.0, 2.5}
	weights := [][]float32{{0.2, 0.8, -0.5, 1.0}, {0.5, -0.91, 0.26, -0.5}, {-0.26, -0.27, 0.17, 0.87}}
	biases := []float32{2.0, 3.0, 0.5}

	layer_outputs := []float32{}
	biasZip, err := goZip.Zip(weights, biases)
	if err != nil {
		fmt.Println(err)
	}

	for _, e := range biasZip {
		nweights, nbias := e.First, e.Second
		var nOutput float32 = 0.0
		inputZip, err := goZip.Zip(inputs, nweights)
		if err != nil {
			fmt.Println(err)
		}
		for _, e := range inputZip {
			nInput, weight := e.First, e.Second
			nOutput += nInput * weight
		}
		nOutput += nbias
		layer_outputs = append(layer_outputs, nOutput)
	}
	fmt.Println(layer_outputs)
	return "Chapter 2"
}
