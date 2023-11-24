package chapter2

import (
	"fmt"

	"github.com/Dorbii/NNFS/internal/utils/printSolution"
	"github.com/Dorbii/NNFS/pkg/configManager"
	t "gorgonia.org/tensor"
)

func Chapter2_02() string {
	conf := configManager.Read("C:\\Users\\Steve\\Documents\\Github\\NNFS\\configs\\chapter2_02.config")

	inputs := t.New(t.WithBacking(conf.Inputs))
	weights := t.New(t.WithShape(3, 4), t.WithBacking(conf.Weights))
	biases := t.New(t.WithBacking(conf.Biases))

	dotProduct, err := t.Dot(weights, inputs)
	if err != nil {
		fmt.Println(err)
	}
	output, err := t.Add(dotProduct, biases)
	if err != nil {
		fmt.Println(err)
	}
	s := printSolution.PrintSolution(printSolution.GetFunctionName(Chapter2_02), output)
	return s
}
