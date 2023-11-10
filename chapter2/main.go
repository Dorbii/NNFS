package main

import "fmt"

func main() {
	inputs := []float32{1.0, 2.0, 3.0, 2.5}
	weights := []float32{0.2, 0.8, -0.5, 1.0}
	var bias float32 = 2

	output := (inputs[0]*weights[0] +
		inputs[1]*weights[1] +
		inputs[2]*weights[2] +
		inputs[3]*weights[3] + bias)
	fmt.Println(output)
}
