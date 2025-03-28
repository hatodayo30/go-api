package handler

import (
	"go-college/internal/interface/dto"
	"go-college/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RankingHandler struct {
	rankingUsecase usecase.RankingUsecase
}

func NewRankingHandler(rankingUsecase usecase.RankingUsecase) *RankingHandler {
	return &RankingHandler{rankingUsecase: rankingUsecase}
}

func (h *RankingHandler) HandleRankingList() echo.HandlerFunc {
	return func(c echo.Context) error {
		start, err := strconv.Atoi(c.QueryParam("start"))
		if err != nil || start < 1 {
			start = 1
		}
		limit := 10

		// ユースケースからランキング取得
		rankings, err := h.rankingUsecase.GetRanking(start, limit)
		if err != nil {
			log.Println("ランキング取得エラー:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ランキングを取得できませんでした"})
		}

		// DTO 変換
		var response []dto.RankingUser
		for i, r := range rankings {
			response = append(response, dto.RankingUser{
				UserID:   r.UserID,
				UserName: r.UserName,
				Rank:     int32(start + i),
				Score:    r.Score,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"ranks": response,
		})
	}
}
