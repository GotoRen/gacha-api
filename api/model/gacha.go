package model

// ガチャの設定
type GachaConfig struct {
	Reality     int     `json:"reality"`
	Probability float64 `json:"probability"`
}

type GroupedCharacters struct {
	Reality     int
	Probability float64
	IDs         []int
}

type GachaDrawRequest struct {
	Count int `json:"count"`
}
