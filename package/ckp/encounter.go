package ckp

type EncounterArea struct {
	Area  string          `json:"area"`
	Pools []EncounterPool `json:"pools"`
}

type EncounterPool struct {
	Type       string
	Encounters []Encounter
	Table      *EncounterTable
	PoolRef    string
}

type EncounterPools struct {
	Fishing   map[string]EncounterFishing `json:"fishing"`
	Headbutt  map[string][]Encounter      `json:"headbutt"`
	RockSmash map[string][]Encounter      `json:"rock"`
}

type EncounterFishing struct {
	Good  EncounterTable `json:"good"`
	Old   EncounterTable `json:"old"`
	Super EncounterTable `json:"super"`
}

type EncounterSpecial struct {
	Pool []Encounter `json:"pool"`
	Type string      `json:"type"`
}

type EncounterTable struct {
	Day     []Encounter `json:"day"`
	Night   []Encounter `json:"night"`
	Morning []Encounter `json:"morning,omitempty"`
}

type Encounter struct {
	Pokemon string `json:"pokemon"`
	Chance  int    `json:"chance"`
	Level   int    `json:"level"`
	Extra   string `json:"extra,omitempty"`
}
