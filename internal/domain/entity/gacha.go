package entity

type CollectionGachaItem struct {
	CollectionID string `json:"collectionID"`
	Name         string `json:"name"`
	Rarity       int    `json:"rarity"`
	Ratio        int    `json:"ratio"`
	IsNew        bool   `json:"isNew"`
}
