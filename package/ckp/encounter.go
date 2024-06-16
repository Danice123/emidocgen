package ckp

import (
	"encoding/json"
	"errors"
)

type TIME_OF_DAY string

const MORNING = TIME_OF_DAY("morning")
const DAY = TIME_OF_DAY("day")
const NIGHT = TIME_OF_DAY("night")

type Encounter struct {
	Area  string          `json:"area"`
	Pools []EncounterPool `json:"pools"`
}

type EncounterPool struct {
	Type       string                             `json:"-"`
	PoolString string                             `json:"-"`
	PoolSlice  []EncounterPokemon                 `json:"-"`
	PoolMap    map[TIME_OF_DAY][]EncounterPokemon `json:"-"`
}

func (ep *EncounterPool) UnmarshalJSON(data []byte) error {
	var ept struct {
		Type string `json:"type"`
	}
	err := json.Unmarshal(data, &ept)
	if err != nil {
		return nil
	}
	switch ept.Type {
	case "old-rod":
		fallthrough
	case "good-rod":
		fallthrough
	case "super-rod":
		fallthrough
	case "headbutt":
		fallthrough
	case "rock":
		var sPool struct {
			Pool string `json:"pool"`
		}
		err := json.Unmarshal(data, &sPool)
		if err != nil {
			return nil
		}
		*ep = EncounterPool{
			Type:       ept.Type,
			PoolString: sPool.Pool,
		}
	case "surfing":
		fallthrough
	case "static":
		fallthrough
	case "trade":
		fallthrough
	case "buy":
		fallthrough
	case "bug-catching-contest":
		fallthrough
	case "gift":
		var slPool struct {
			Pool []EncounterPokemon `json:"pool"`
		}
		err := json.Unmarshal(data, &slPool)
		if err != nil {
			return nil
		}
		*ep = EncounterPool{
			Type:      ept.Type,
			PoolSlice: slPool.Pool,
		}
	case "swarm":
		fallthrough
	case "swarm-old-rod":
		fallthrough
	case "swarm-good-rod":
		fallthrough
	case "swarm-super-rod":
		fallthrough
	case "walking":
		var mPool struct {
			Pool map[TIME_OF_DAY][]EncounterPokemon `json:"pool"`
		}
		err := json.Unmarshal(data, &mPool)
		if err != nil {
			return nil
		}
		*ep = EncounterPool{
			Type:    ept.Type,
			PoolMap: mPool.Pool,
		}
	default:
		return errors.New("Unsupported type: " + ept.Type)
	}
	return nil
}

func (ep *EncounterPool) MarshalJSON() ([]byte, error) {
	switch ep.Type {
	case "old-rod":
		fallthrough
	case "good-rod":
		fallthrough
	case "super-rod":
		fallthrough
	case "headbutt":
		fallthrough
	case "rock":
		stringEP := struct {
			Type string `json:"type"`
			Pool string `json:"pool"`
		}{
			Type: ep.Type,
			Pool: ep.PoolString,
		}
		return json.Marshal(stringEP)
	case "surfing":
		fallthrough
	case "static":
		fallthrough
	case "trade":
		fallthrough
	case "buy":
		fallthrough
	case "bug-catching-contest":
		fallthrough
	case "gift":
		slEP := struct {
			Type string             `json:"type"`
			Pool []EncounterPokemon `json:"pool"`
		}{
			Type: ep.Type,
			Pool: ep.PoolSlice,
		}
		return json.Marshal(slEP)
	case "swarm":
		fallthrough
	case "swarm-old-rod":
		fallthrough
	case "swarm-good-rod":
		fallthrough
	case "swarm-super-rod":
		fallthrough
	case "walking":
		mEP := struct {
			Type string                             `json:"type"`
			Pool map[TIME_OF_DAY][]EncounterPokemon `json:"pool"`
		}{
			Type: ep.Type,
			Pool: ep.PoolMap,
		}
		return json.Marshal(mEP)
	default:
		return nil, errors.New("Unsupported type: " + ep.Type)
	}
}

type EncounterPokemon struct {
	Chance  int    `json:"chance"`
	Level   int    `json:"level"`
	Pokemon string `json:"pokemon"`
}

type EncounterPools struct {
	Fishing  map[string]FishingPools       `json:"fishing"`
	Headbutt map[string][]EncounterPokemon `json:"headbutt"`
	Rock     map[string][]EncounterPokemon `json:"rock"`
}

type FishingPools struct {
	Old   map[TIME_OF_DAY][]EncounterPokemon `json:"old"`
	Good  map[TIME_OF_DAY][]EncounterPokemon `json:"good"`
	Super map[TIME_OF_DAY][]EncounterPokemon `json:"super"`
}
