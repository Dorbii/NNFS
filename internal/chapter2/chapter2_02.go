package chapter2

import (
	"fmt"

	configManager "github.com/Dorbii/NNFS/pkg/configManager"

	t "gorgonia.org/tensor"
)

func chapter2_02() {
	conf := configManager.Read("../../configs/config.json")
	inputs := t.New(t.WithBacking(conf.Inputs))
	fmt.Println(inputs)
	//layerOutputs := t.Dot(t.New(inputs), t.New(weights))
}
