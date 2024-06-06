package transforms

import (
	t "gorgonia.org/tensor"
)

func Bias(newShape t.Shape, values []float64) *t.Dense {
	newBacking := make([]float64, newShape[0]*newShape[1])
	for i := range newBacking {
		newBacking[i] = values[i%len(values)]
	}
	return t.New(t.WithShape(newShape...), t.WithBacking(newBacking))
}
