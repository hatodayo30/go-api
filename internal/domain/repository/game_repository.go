package repository

import "database/sql"

type FinishRepository interface {
	UpdateUserScoreWithTx(tx *sql.Tx, userID string, score int) error
	UpdateUserCoinsWithTx(tx *sql.Tx, userID string, coins int) error
}
