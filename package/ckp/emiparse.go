package ckp

import (
	"encoding/json"
	"os"
)

type TypeMatchup struct {
	Attacker   PokeType `json:"attacker"`
	Defender   PokeType `json:"defender"`
	Multiplier float32  `json:"multiplier"`
}

type EmiCalcData struct {
	Encounters   []interface{} `json:"encounters"`
	Pools        interface{}   `json:"encounter_pools"`
	Items        []Item        `json:"items"`
	Landmarks    []Landmark    `json:"landmarks"`
	Pokemon      []Pokemon     `json:"pokemon"`
	Moves        []PokeMove    `json:"moves"`
	Trainers     []Trainer     `json:"trainers"`
	TypeMatchups []TypeMatchup `json:"type_matchups"`

	PokeMap map[string]int `json:"-"`
}

func ParseEmiData(file string) (*EmiCalcData, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var data EmiCalcData
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	data.PokeMap = map[string]int{}
	for _, p := range data.Pokemon {
		data.PokeMap[p.Name] = p.Pokedex
	}

	return &data, nil
}
