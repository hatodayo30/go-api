package dto

type GameFinishRequest struct {
	Score int32 `json:"score"`
}

type GameFinishResponse struct {
	Coin int `json:"coin"`
}
