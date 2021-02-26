package models

type Area struct {
	PosX  int `json:"posX"`
	PosY  int `json:"posY"`
	SizeX int `json:"sizeX"`
	SizeY int `json:"sizeY"`
}

type ExploreResp struct {
	Area   Area `json:"area"`
	Amount int  `json:"amount"`
}
