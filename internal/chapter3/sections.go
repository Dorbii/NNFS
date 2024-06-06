package chapter3

import (
	"fmt"
	"slices"

	"github.com/Dorbii/NNFS/internal/datasets"
	"github.com/Dorbii/NNFS/internal/transforms"
	"github.com/Dorbii/NNFS/internal/utils"
	"github.com/charmbracelet/lipgloss"
	t "gorgonia.org/tensor"
)

type DenseLayer struct {
	Weights t.Tensor
	Biases  []float64
	Output  t.Tensor
}

func Sections() {
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("14")).SetString("Chapter 3"))
	utils.Header("Dense Layer:")
	dLayer()
}

func dLayer() {
	d := &DenseLayer{}
	x, _ := datasets.SpiralData(100, 3)
	d.New(2, 3)
	d.Forward(x)
	utils.Output(d.Output.Slice(t.S(0, 5), t.S(0, 5)))
}

func (d *DenseLayer) New(nInputs, nNeurons int) {
	weights := t.New(t.WithShape(nInputs, nNeurons),
		t.WithBacking(t.Random(t.Float64, nInputs*nNeurons)))
	weights, err := weights.MulScalar(0.01, false)
	if err != nil {
		utils.Error("mulScalar", weights)
	}
	biases := make([]float64, nNeurons)
	d.Weights = weights
	d.Biases = biases
}

func (d *DenseLayer) Forward(fInputs [][]float64) {
	inputs := t.New(t.WithShape(len(fInputs), len(fInputs[0])), t.WithBacking(slices.Concat(fInputs...)))
	dotProduct, _ := t.Dot(inputs, d.Weights)
	biases := transforms.Bias(dotProduct.Shape(), d.Biases)
	v, _ := t.Add(dotProduct, biases)
	d.Output = v
}
