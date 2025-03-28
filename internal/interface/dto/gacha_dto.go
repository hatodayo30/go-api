package dto

import "go-college/internal/domain/entity"

type GachaDrawRequest struct {
	Times int `json:"times"`
}

type GachaDrawResponse struct {
	Results []entity.CollectionGachaItem `json:"results"`
}
