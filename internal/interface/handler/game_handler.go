package handler

import (
	"go-college/internal/context/auth"
	"go-college/internal/infrastructure/db"
	"go-college/internal/interface/dto"
	"go-college/internal/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GameHandler struct {
	gameUsecase usecase.GameUsecase
}

func NewGameHandler(gameUsecase usecase.GameUsecase) *GameHandler {
	return &GameHandler{gameUsecase: gameUsecase}
}

func (h *GameHandler) HandleGameFinish() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.GameFinishRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		ctx := c.Request().Context()
		userID, ok := auth.GetUserIDFromContext(ctx)
		if !ok || userID == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}

		tx, err := db.Conn.Begin()
		if err != nil {
			log.Println("トランザクション開始失敗:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "transaction failed"})
		}

		reward, err := h.gameUsecase.FinishGame(tx, userID, req.Score)
		if err != nil {
			tx.Rollback()
			log.Println("ゲーム終了処理失敗:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if err := tx.Commit(); err != nil {
			log.Println("コミット失敗:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "commit failed"})
		}

		return c.JSON(http.StatusOK, dto.GameFinishResponse{Coin: reward})
	}
}
