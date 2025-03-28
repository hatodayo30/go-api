package repository

import (
	"go-college/internal/domain/entity"
)

//go:generate mockgen -source=collection.go -destination=./mock/collection_mock.go -package=mock
type CollectionRepository interface {
	GetUserCollectionList(userID string) ([]entity.CollectionItem, int, int, error)
}
