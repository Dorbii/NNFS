package configManager

import (
	"encoding/json"
	"os"
)

func read(file string, c interface{}) {
	f, _ := os.Open(file)
	json.NewDecoder(f).Decode(&c)
}

// func Save(fileName string) {
//If dynamic config is needed
// }
