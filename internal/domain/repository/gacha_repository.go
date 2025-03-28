package repository

import (
	"go-college/internal/domain/entity"
)

// GachaRepository はガチャに関連するデータアクセスのインターフェース
type GachaRepository interface {
	GetUserCoins(userID string) (int, error)
	UpdateUserCoins(userID string, amount int) error
	GetAllGachaItems() ([]entity.CollectionGachaItem, error)
	GetUserOwnedItems(userID string, itemIDs []string) (map[string]bool, error)
	InsertNewItems(userID string, itemIDs []string) error
}
