package ckp

type Encounter struct {
	Area  string        `json:"area"`
	Pools []interface{} `json:"pools"`
}
