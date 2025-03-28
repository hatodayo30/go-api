package repository

import (
	"database/sql"
	"go-college/internal/domain/repository"
)

type finishRepository struct {
	db *sql.DB
}

func NewFinishRepository(db *sql.DB) repository.FinishRepository {
	return &finishRepository{db: db}
}

func (r *finishRepository) UpdateUserScoreWithTx(tx *sql.Tx, userID string, score int) error {
	var HighScore int
	err := tx.QueryRow("SELECT high_score FROM user WHERE id = ?", userID).Scan(&HighScore)
	if err != nil {
		return err
	}
	if score > HighScore {
		_, err := tx.Exec("UPDATE user SET high_score = ? WHERE id = ?", score, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *finishRepository) UpdateUserCoinsWithTx(tx *sql.Tx, userID string, coins int) error {
	_, err := tx.Exec("UPDATE user SET coin = coin + ? WHERE id = ?", coins, userID)
	return err
}
