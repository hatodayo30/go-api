package usecase

import (
	"errors"
	"math/rand"
	"time"

	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"
)

const (
	GachaCostPerTry = 100
)

var (
	ErrInsufficientCoins = errors.New("コインが不足しています")
)

type GachaUsecase interface {
	ExecuteGacha(userID string, times int) ([]entity.CollectionGachaItem, error)
}

type gachaUsecase struct {
	repo repository.GachaRepository
}

func NewGachaUsecase(repo repository.GachaRepository) GachaUsecase {
	return &gachaUsecase{repo: repo}
}

func (u *gachaUsecase) ExecuteGacha(userID string, times int) ([]entity.CollectionGachaItem, error) {
	cost := times * GachaCostPerTry

	// コイン残高確認
	coins, err := u.repo.GetUserCoins(userID)
	if err != nil {
		return nil, err
	}
	if coins < cost {
		return nil, ErrInsufficientCoins
	}

	// ガチャアイテム取得
	items, err := u.repo.GetAllGachaItems()
	if err != nil {
		return nil, err
	}

	// 抽選実行
	selected := u.draw(items, times)

	// 所持チェックと新規アイテム抽出
	ids := u.extractIDs(selected)
	owned, err := u.repo.GetUserOwnedItems(userID, ids)
	if err != nil {
		return nil, err
	}

	var newItems []string
	for i, item := range selected {
		if owned[item.CollectionID] {
			selected[i].IsNew = false
		} else {
			selected[i].IsNew = true
			newItems = append(newItems, item.CollectionID)
		}
	}

	if len(newItems) > 0 {
		if err := u.repo.InsertNewItems(userID, newItems); err != nil {
			return nil, err
		}
	}

	// コイン減算
	if err := u.repo.UpdateUserCoins(userID, cost); err != nil {
		return nil, err
	}

	return selected, nil
}

// draw はガチャ抽選処理を行う
func (u *gachaUsecase) draw(items []entity.CollectionGachaItem, times int) []entity.CollectionGachaItem {
	var results []entity.CollectionGachaItem
	total := 0
	for _, item := range items {
		total += item.Ratio
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		r := rand.Intn(total)
		acc := 0
		for _, item := range items {
			acc += item.Ratio
			if r < acc {
				results = append(results, item)
				break
			}
		}
	}

	return results
}

// extractIDs はアイテムIDリストを抽出する
func (u *gachaUsecase) extractIDs(items []entity.CollectionGachaItem) []string {
	var ids []string
	for _, item := range items {
		ids = append(ids, item.CollectionID)
	}
	return ids
}
