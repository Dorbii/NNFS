package chapter2

import (
	"fmt"
	"os"

	"github.com/Dorbii/NNFS/internal/transforms"
	"github.com/Dorbii/NNFS/internal/utils"
	"github.com/Dorbii/NNFS/pkg/configManager"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	t "gorgonia.org/tensor"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout)
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR: ").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("204"))
	// Add a custom style for key `err`
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("204"))
	logger.SetStyles(styles)
}
func Sections() {
	config := configManager.ReadYAML("C:\\Users\\Steve\\Documents\\Github\\NNFS\\configs\\chapter2.yaml")
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("14")).SetString("Chapter 2"))
	single(config["single"])
	layer(config["layer"])
	batch(config["batch"])
}

func single(conf configManager.YConfig) {
	utils.Header("Single:")
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
	utils.Output(output)
}

func layer(conf configManager.YConfig) {
	utils.Header("Layer:")
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
	utils.Output(output)

}

func batch(conf configManager.YConfig) {
	utils.Header("Batch:")
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
	utils.Output(output)
}
