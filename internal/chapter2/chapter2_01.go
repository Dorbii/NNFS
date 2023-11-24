package chapter2

import (
	"fmt"

	"github.com/Dorbii/NNFS/pkg/goZip"
)

func Chapter2_01() string {
	inputs := []float32{1.0, 2.0, 3.0, 2.5}
	weights := [][]float32{{0.2, 0.8, -0.5, 1.0}, {0.5, -0.91, 0.26, -0.5}, {-0.26, -0.27, 0.17, 0.87}}
	biases := []float32{2.0, 3.0, 0.5}

	layerOutputs := []float32{}

	weightsBiasZip, err := goZip.Zip(weights, biases)
	if err != nil {
		fmt.Println(err)
	}

	for _, wb := range weightsBiasZip {
		var nOutput float32 = 0.0
		inputWeightZip, err := goZip.Zip(inputs, wb.First) //zip inputs and weights
		if err != nil {
			fmt.Println(err)
		}
		for _, iw := range inputWeightZip {
			nOutput += iw.First * iw.Second //input * weight
		}
		nOutput += wb.Second //bias
		layerOutputs = append(layerOutputs, nOutput)
	}
	fmt.Println(layerOutputs)
	return "Chapter 2"
}
