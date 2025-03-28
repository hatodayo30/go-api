package entity

// DBから取得するランキングデータの構造体
type DBRankingUser struct {
	UserID   string `db:"id"`
	UserName string `db:"name"`
	Rank     int32  `db:"rank"`
	Score    int32  `db:"high_score"`
}
