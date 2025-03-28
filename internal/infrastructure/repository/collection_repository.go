package repository

import (
	"database/sql"
	"fmt"
	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"
	"strings"
)

// collectionRepository 構造体 (インターフェースの実装)
type collectionRepository struct {
	db *sql.DB
}

// NewCollectionRepository CollectionRepository のインスタンスを生成
func NewCollectionRepository(database *sql.DB) repository.CollectionRepository {
	return &collectionRepository{db: database}
}

// GetUserCollectionList ユーザーのコレクションリストを取得
func (r *collectionRepository) GetUserCollectionList(userID string) ([]entity.CollectionItem, int, int, error) {
	// **全コレクション数を取得**
	var totalCollections int
	err := r.db.QueryRow("SELECT COUNT(*) FROM collection_item").Scan(&totalCollections)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("コレクション数取得エラー: %v", err)
	}

	// **全アイテムのIDリストを取得**
	rows, err := r.db.Query("SELECT id, name, rarity FROM collection_item")
	if err != nil {
		return nil, 0, 0, fmt.Errorf("コレクションアイテム取得エラー: %v", err)
	}
	defer rows.Close()

	// **アイテムリストとIDリストを作成**
	allItems := make(map[string]entity.CollectionItem)
	var itemIDs []string

	for rows.Next() {
		var item entity.CollectionItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Rarity); err != nil {
			return nil, 0, 0, fmt.Errorf("アイテムスキャンエラー: %v", err)
		}
		allItems[item.ID] = item
		itemIDs = append(itemIDs, item.ID)
	}

	// **所持アイテムの判定**
	ownedCollections := 0

	if len(itemIDs) > 0 {
		placeholders := strings.Repeat("?,", len(itemIDs))
		placeholders = placeholders[:len(placeholders)-1]

		query := fmt.Sprintf("SELECT collection_item_id FROM user_collection_item WHERE user_id = ? AND collection_item_id IN (%s)", placeholders)

		args := make([]interface{}, 0, len(itemIDs)+1)
		args = append(args, userID)
		for _, id := range itemIDs {
			args = append(args, id)
		}

		ownedRows, err := r.db.Query(query, args...)
		if err != nil {
			return nil, 0, 0, fmt.Errorf("所持アイテム取得エラー: %v", err)
		}
		defer ownedRows.Close()

		for ownedRows.Next() {
			var collectionID string
			if err := ownedRows.Scan(&collectionID); err != nil {
				return nil, 0, 0, fmt.Errorf("所持アイテムスキャンエラー: %v", err)
			}
			if item, exists := allItems[collectionID]; exists {
				item.HasItem = true
				allItems[collectionID] = item
				ownedCollections++
			}
		}
	}

	items := make([]entity.CollectionItem, 0, len(allItems))
	for _, item := range allItems {
		items = append(items, item)
	}

	return items, ownedCollections, totalCollections, nil
}
