package dto

import "go-college/internal/domain/entity"

// CollectionResponse はAPIのレスポンス全体を表す構造体
type CollectionResponse struct {
	Collections      []entity.CollectionItem `json:"collections"`
	OwnedCollections int                     `json:"ownedCollections"`
	TotalCollections int                     `json:"totalCollections"`
}
