package usecase

import (
	"database/sql"
	"errors"
	"go-college/internal/domain/repository"
)

type GameUsecase interface {
	FinishGame(tx *sql.Tx, userID string, score int32) (int, error)
}

type gameUsecase struct {
	repo repository.FinishRepository
}

func NewGameUsecase(repo repository.FinishRepository) GameUsecase {
	return &gameUsecase{repo: repo}
}

func (u *gameUsecase) FinishGame(tx *sql.Tx, userID string, score int32) (int, error) {
	if score < 0 {
		return 0, errors.New("0以上のスコアを入力してください")
	}

	reward := calculateReward(score)

	if err := u.repo.UpdateUserScoreWithTx(tx, userID, int(score)); err != nil {
		return 0, err
	}
	if err := u.repo.UpdateUserCoinsWithTx(tx, userID, reward); err != nil {
		return 0, err
	}

	return reward, nil
}

func calculateReward(score int32) int {
	return int(score) / 10
}
