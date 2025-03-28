package repository

import "go-college/internal/domain/entity"

type RankingRepository interface {
	RankingList(start int, limit int) ([]entity.DBRankingUser, error)
}
