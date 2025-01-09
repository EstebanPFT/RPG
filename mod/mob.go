package mod

import (
	"encoding/json"
	"fmt"
	"os"
)

type Mob struct {
	Name   string `json:"Name"`
	PV     int    `json:"PV"`
	MaxPV  int    `json:"MaxPV"`
	Force  int    `json:"Force"`
}


func Arrivagedemob(random string) (Mob, error) {
	file, err := os.Open("./data/Mob.json")
	if err != nil {
		return Mob{}, fmt.Errorf("err: %v", err)
	}
	defer file.Close()

	var mobs map[string]Mob
	err = json.NewDecoder(file).Decode(&mobs)
	if err != nil {
		return Mob{}, fmt.Errorf("err json : %v", err)
	}

	choosedMob, exists := mobs[random]
	if !exists {
		return Mob{}, fmt.Errorf("mob '%s' pas trouver", random)
	}

	return choosedMob, nil
}

