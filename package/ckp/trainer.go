package ckp

type Trainer struct {
	Area  string        `json:"area"`
	Name  string        `json:"name"`
	Team  []interface{} `json:"team"`
	Meta  string        `json:"meta,omitempty"`
	B2b   bool          `json:"b2b,omitempty"`
	Order int           `json:"order,omitempty"`
}

type TrainerPokemon struct {
	Name  string       `json:"name"`
	Level int          `json:"level"`
	Item  string       `json:"item,omitempty"`
	Moves []string     `json:"moves"`
	Dvs   map[Stat]int `json:"dvs,omitempty"`
}
