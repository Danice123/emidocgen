package ckp

type Stat string

const Atk = Stat("atk")
const Def = Stat("def")
const HP = Stat("hp")
const SpA = Stat("spa")
const SpD = Stat("spd")
const Spe = Stat("spe")

type PokeType string

const Normal = PokeType("normal")
const Fire = PokeType("fire")
const Water = PokeType("water")
const Electric = PokeType("electric")
const Grass = PokeType("grass")
const Ice = PokeType("ice")
const Fighting = PokeType("fighting")
const Poison = PokeType("poison")
const Ground = PokeType("ground")
const Flying = PokeType("flying")
const Psychic = PokeType("psychic")
const Bug = PokeType("bug")
const Rock = PokeType("rock")
const Ghost = PokeType("ghost")
const Dragon = PokeType("dragon")
const Dark = PokeType("dark")
const Steel = PokeType("steel")

type Pokemon struct {
	Name       string       `json:"name"`
	Pokedex    int          `json:"pokedex"`
	BaseExp    int          `json:"base_experience"`
	Evolutions []Evolution  `json:"evolutions,omitempty"`
	Gender     string       `json:"gender"`
	Items      []ItemHold   `json:"items"`
	Learnset   []Learnset   `json:"learnset"`
	Stats      map[Stat]int `json:"stats"`
	TmHm       []string     `json:"tmhm"`
	Types      []PokeType   `json:"types"`
}

type Evolution struct {
	Target string `json:"into"`
	Level  int    `json:"level,omitempty"`
	Method string `json:"method"`
	Item   string `json:"item,omitempty"`
}

type ItemHold struct {
	Chance float32 `json:"chance"`
	Item   string  `json:"item"`
}

type Learnset struct {
	Level int    `json:"level"`
	Move  string `json:"move"`
}
