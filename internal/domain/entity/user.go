package entity

// User userテーブルデータ
type User struct {
	ID        string
	AuthToken string
	Name      string
	HighScore int32
	Coin      int32
}
