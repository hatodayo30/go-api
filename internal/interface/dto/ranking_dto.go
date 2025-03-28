package dto

// ランキングユーザー構造体
type RankingUser struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Rank     int32  `json:"rank"`
	Score    int32  `json:"score"`
}
