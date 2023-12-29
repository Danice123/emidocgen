package ckp

type PokeMove struct {
	Name     string             `json:"name"`
	Index    int                `json:"index,omitempty"`
	Power    int                `json:"power"`
	PP       int                `json:"pp"`
	Type     PokeType           `json:"type"`
	Accuracy int                `json:"accuracy"`
	Effects  map[string]float32 `json:"effects,omitempty"`
	Extra    []string           `json:"extra,omitempty"`
}
