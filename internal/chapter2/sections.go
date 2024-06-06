package chapter2

import (
	"fmt"

	"github.com/Dorbii/NNFS/internal/transforms"
	"github.com/Dorbii/NNFS/internal/utils"
	"github.com/Dorbii/NNFS/pkg/configManager"
	"github.com/charmbracelet/lipgloss"
	t "gorgonia.org/tensor"
)

func Sections() {
	config := configManager.ReadYAML("C:\\Users\\Steve\\Documents\\Github\\NNFS\\configs\\chapter2.yaml")
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("14")).SetString("Chapter 2"))
	utils.Header("Single:")
	singleOutput := single(config["single"])
	utils.Output(singleOutput)
	utils.Header("Layer:")
	layerOutput := layer(config["layer"])
	utils.Output(layerOutput)
	utils.Header("Batch:")
	batchOutput := batch(config["batch"])
	utils.Output(batchOutput)
	utils.Header("HiddenLater")
	hiddenLayerOutput := hiddenLayer(config["hiddenLayer"])
	utils.Output(hiddenLayerOutput)
}

func single(conf configManager.YConfig) t.Tensor {
	inputs := t.New(t.WithBacking(conf.Inputs))
	weights := t.New(t.WithBacking(conf.Weights))
	biases := t.New(t.WithBacking(conf.Biases))

	dotProduct, err := t.Dot(inputs, weights)
	if err != nil {
		utils.Error("dotProduct", inputs, weights)
	}
	output, err := t.Add(dotProduct, biases)
	if err != nil {
		utils.Error("sum", dotProduct, biases)
	}
	return output
}

func layer(conf configManager.YConfig) t.Tensor {
	inputs := t.New(t.WithBacking(conf.Inputs))
	weights := t.New(t.WithShape(3, 4), t.WithBacking(conf.Weights))
	biases := t.New(t.WithBacking(conf.Biases))

	dotProduct, err := t.Dot(weights, inputs)
	if err != nil {
		utils.Error("dotProduct", weights, inputs)
	}
	output, err := t.Add(dotProduct, biases)
	if err != nil {
		utils.Error("sum", dotProduct, biases)
	}
	return output

}

func batch(conf configManager.YConfig) t.Tensor {
	inputs := t.New(t.WithShape(3, 4), t.WithBacking(conf.Inputs))
	weights := t.New(t.WithShape(3, 4), t.WithBacking(conf.Weights))
	weights.T()

	dotProduct, err := t.Dot(inputs, weights)
	if err != nil {
		utils.Error("dotProduct", inputs, weights)
	}
	biases := transforms.Bias(dotProduct.Shape(), conf.Biases)
	output, err := t.Add(dotProduct, biases)
	if err != nil {
		utils.Error("sum", dotProduct, biases)
	}
	return output
}

func hiddenLayer(conf configManager.YConfig) t.Tensor {
	weights2 := t.New(t.WithShape(3, 3), t.WithBacking(conf.Weights2))
	weights2.T()
	batchOutput := batch(conf)
	dotProduct, err := t.Dot(batchOutput, weights2)
	if err != nil {
		utils.Error("dotProduct", batchOutput, weights2)
	}
	biases := transforms.Bias(dotProduct.Shape(), conf.Biases2)
	output, err := t.Add(dotProduct, biases)
	if err != nil {
		utils.Error("sum", dotProduct, biases)
	}
	return output

}
