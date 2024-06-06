package chapter4

import (
	"fmt"

	"github.com/Dorbii/NNFS/internal/chapter3"
	"github.com/Dorbii/NNFS/internal/datasets"
	"github.com/Dorbii/NNFS/internal/utils"
	"github.com/Dorbii/NNFS/pkg/configManager"
	"github.com/charmbracelet/lipgloss"
	t "gorgonia.org/tensor"
)

type ActivationReLu struct {
	MinVal float64
	Output t.Tensor
}

func Sections() {
	config := configManager.ReadYAML("C:\\Users\\Steve\\Documents\\Github\\NNFS\\configs\\chapter4.yaml")
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("14")).SetString("Chapter 4"))
	utils.Header("ReLuActivation:")
	reluOutput := reluActivation(config["reluActivation"])
	utils.Output(reluOutput)

	utils.Header("DenseLayer Activation:")
	denseLayerActivation()
}

func reluActivation(conf configManager.YConfig) []float64 {
	output := []float64{}
	for i := range conf.Inputs {
		output = append(output, max(0.0, conf.Inputs[i]))
	}
	return output
}

func denseLayerActivation() {
	d := &chapter3.DenseLayer{}
	a := &ActivationReLu{MinVal: 0.0}
	x, _ := datasets.SpiralData(100, 3)
	d.New(2, 3)
	d.Forward(x)
	a.Forward(d.Output)
	utils.Output(a.Output.Slice(t.S(0, 5), t.S(0, 5)))
}

func (a *ActivationReLu) Forward(inputs t.Tensor) {
	data := t.Dataer(inputs).Data().([]float64)
	output := t.New(t.WithShape(inputs.Shape()...), t.WithBacking(activationMax(a.MinVal, data)))
	a.Output = output
}

func activationMax(minVal float64, input []float64) []float64 {
	output := []float64{}
	for i := range input {
		output = append(output, max(minVal, input[i]))
	}
	return output
}
