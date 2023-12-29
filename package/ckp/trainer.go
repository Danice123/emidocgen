package ckp

type Trainer struct {
	Area  string        `json:"area"`
	Name  string        `json:"name"`
	Team  []interface{} `json:"team"`
	Order int           `json:"order"`
}

type TrainerPokemon struct {
	Name  string       `json:"name"`
	Level int          `json:"level"`
	Item  string       `json:"item,omitempty"`
	Moves []string     `json:"moves"`
	Dvs   map[Stat]int `json:"dvs,omitempty"`
}
