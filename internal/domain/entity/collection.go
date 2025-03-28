package entity

// CollectionItem はコレクションアイテムのデータ構造
type CollectionItem struct {
	ID      string `json:"collectionID" db:"id"`
	Name    string `json:"name" db:"name"`
	Rarity  int32  `json:"rarity" db:"rarity"`
	HasItem bool   `json:"hasItem"`
}
