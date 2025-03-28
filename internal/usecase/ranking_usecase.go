package usecase

import (
	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"
)

type RankingUsecase interface {
	GetRanking(start, limit int) ([]entity.DBRankingUser, error)
}

type rankingUsecase struct {
	rankingRepo repository.RankingRepository
}

func NewRankingUsecase(rankingRepo repository.RankingRepository) RankingUsecase {
	return &rankingUsecase{rankingRepo: rankingRepo}
}

func (u *rankingUsecase) GetRanking(start, limit int) ([]entity.DBRankingUser, error) {
	if start < 1 {
		start = 1
	}
	return u.rankingRepo.RankingList(start, limit)
}
