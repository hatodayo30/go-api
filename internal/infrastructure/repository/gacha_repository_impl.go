package repository

import (
	"database/sql"
	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"
	"strings"
)

type gachaRepositoryImpl struct {
	db *sql.DB
}

func NewGachaRepository(db *sql.DB) repository.GachaRepository {
	return &gachaRepositoryImpl{db: db}
}

func (r *gachaRepositoryImpl) GetUserCoins(userID string) (int, error) {
	var coins int
	err := r.db.QueryRow("SELECT coin FROM user WHERE id = ?", userID).Scan(&coins)
	return coins, err
}

func (r *gachaRepositoryImpl) UpdateUserCoins(userID string, amount int) error {
	_, err := r.db.Exec("UPDATE user SET coin = coin - ? WHERE id = ?", amount, userID)
	return err
}

func (r *gachaRepositoryImpl) GetAllGachaItems() ([]entity.CollectionGachaItem, error) {
	query := `
        SELECT c.id, c.name, c.rarity, g.ratio
        FROM collection_item c
        JOIN gacha_probability g ON c.id = g.collection_item_id
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entity.CollectionGachaItem
	for rows.Next() {
		var item entity.CollectionGachaItem
		if err := rows.Scan(&item.CollectionID, &item.Name, &item.Rarity, &item.Ratio); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *gachaRepositoryImpl) GetUserOwnedItems(userID string, itemIDs []string) (map[string]bool, error) {
	if len(itemIDs) == 0 {
		return map[string]bool{}, nil
	}

	placeholders := strings.Repeat("?,", len(itemIDs)-1) + "?"
	query := `
        SELECT collection_item_id FROM user_collection_item 
        WHERE user_id = ? AND collection_item_id IN (` + placeholders + `)
    `

	args := make([]interface{}, len(itemIDs)+1)
	args[0] = userID
	for i, id := range itemIDs {
		args[i+1] = id
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ownedItems := make(map[string]bool)
	for rows.Next() {
		var collectionID string
		if err := rows.Scan(&collectionID); err != nil {
			return nil, err
		}
		ownedItems[collectionID] = true
	}
	return ownedItems, nil
}

func (r *gachaRepositoryImpl) InsertNewItems(userID string, itemIDs []string) error {
	if len(itemIDs) == 0 {
		return nil
	}

	values := strings.Repeat("(?, ?),", len(itemIDs)-1) + "(?, ?)"
	query := `
        INSERT INTO user_collection_item (user_id, collection_item_id)
        VALUES ` + values

	args := make([]interface{}, len(itemIDs)*2)
	for i, id := range itemIDs {
		args[i*2] = userID
		args[i*2+1] = id
	}

	_, err := r.db.Exec(query, args...)
	return err
}
