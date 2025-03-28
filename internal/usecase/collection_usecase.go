package usecase

import (
	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"
)

type CollectionUsecase interface {
	GetUserCollectionList(userID string) ([]entity.CollectionItem, int, int, error)
}

type collectionUsecase struct {
	collectionRepo repository.CollectionRepository
}

func NewCollectionUsecase(collectionRepo repository.CollectionRepository) CollectionUsecase {
	return &collectionUsecase{collectionRepo: collectionRepo}
}

func (u *collectionUsecase) GetUserCollectionList(userID string) ([]entity.CollectionItem, int, int, error) {
	return u.collectionRepo.GetUserCollectionList(userID)
}
