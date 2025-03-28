package repository

import (
	"database/sql"
	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"
	"go-college/internal/infrastructure/db"
)

type rankingRepository struct {
	db *sql.DB
}

func NewRankingRepository(db *sql.DB) repository.RankingRepository {
	return &rankingRepository{db: db}
}

// ランキング取得（start位から `limit` 件取得）
func (r *rankingRepository) RankingList(start int, limit int) ([]entity.DBRankingUser, error) {
	query := `
		SELECT id, name, high_score
		FROM user
		ORDER BY high_score DESC, id ASC
		LIMIT ? OFFSET ?;
	`
	rows, err := db.Conn.Query(query, limit, start-1) // OFFSET は 0-index
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rankings []entity.DBRankingUser
	for rows.Next() {
		var user entity.DBRankingUser
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Score); err != nil {
			return nil, err
		}
		user.Rank = 0
		rankings = append(rankings, user)
	}

	return rankings, nil
}
