package ckp

type Landmark struct {
	Id        int           `json:"id"`
	Locations []string      `json:"locations"`
	Name      string        `json:"name"`
	Items     []MapItem     `json:"items"`
	Positions []MapPosition `json:"positions"`
}

type MapPosition struct {
	X      int `json:"x,omitempty"`
	Y      int `json:"y,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type MapItem struct {
	Item   string `json:"item"`
	Amount string `json:"amount"`
	Info   string `json:"info"`
}
