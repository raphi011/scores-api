package dtos

type Statistic struct {
	Points int `json:"points"`
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
}
