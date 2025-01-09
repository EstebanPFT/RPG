package mod

import (
	"encoding/json"
	"fmt"
	"os"
)

type Boss struct {
	Name   string `json:"Name"`
	PV     int    `json:"PV"`
	MaxPV  int    `json:"MaxPV"`
	Force  int    `json:"Force"`
}

func Arrivagedeboss(random2 string) (Boss, error) {
	file, err := os.Open("./data/Boss.json") 
	if err != nil {
		return Boss{}, fmt.Errorf("erreur lors de l'ouverture du fichier : %v", err)
	}
	defer file.Close()

	var bossDataGen map[string]Boss
	err = json.NewDecoder(file).Decode(&bossDataGen)
	if err != nil {
		return Boss{}, fmt.Errorf("erreur lors du décodage JSON : %v", err)
	}

	choosedBoss, exists := bossDataGen[random2]
	if !exists {
		return Boss{}, fmt.Errorf("boss '%s' non trouvé", random2)
	}

	return choosedBoss, nil
}
