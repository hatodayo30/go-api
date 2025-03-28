package handler

import (
	"go-college/internal/constant"
	"go-college/internal/context/auth"
	"go-college/internal/interface/dto"
	"go-college/internal/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func (h *UserHandler) HandleSettingGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.SettingGetResponse{
			GachaCoinConsumption: constant.GachaCoinConsumption,
		})
	}
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

// ユーザー作成
func (h *UserHandler) HandleUserCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.UserCreateRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		user, err := h.userUsecase.CreateUser(req.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
		}

		return c.JSON(http.StatusCreated, dto.UserGetResponse{
			ID:        user.ID,
			Name:      user.Name,
			HighScore: user.HighScore,
			Coin:      user.Coin,
		})
	}
}

// ユーザー取得
func (h *UserHandler) HandleUserGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID, ok := auth.GetUserIDFromContext(ctx)
		if !ok || userID == "" {
			log.Println("ユーザーIDが見つかりません")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "ユーザー認証が必要です"})
		}
		user, err := h.userUsecase.GetUserByID(userID)
		if err != nil {
			log.Println("ユーザー取得エラー", err)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "ユーザーが見つかりません"})
		}
		return c.JSON(http.StatusOK, dto.UserGetResponse{
			ID:        user.ID,
			Name:      user.Name,
			HighScore: user.HighScore,
			Coin:      user.Coin,
		})
	}
}

// ユーザー更新処理
func (h *UserHandler) HandleUserUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID, ok := auth.GetUserIDFromContext(ctx)
		if !ok || userID == "" {
			log.Println("ユーザーIDが見つかりません")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "ユーザー認証が必要です"})
		}
		req := new(dto.UserUpdateRequest)
		if err := c.Bind(req); err != nil {
			log.Println("リクエストバインドエラー:", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		err := h.userUsecase.UpdateUser(userID, req.Name)
		if err != nil {
			log.Println("ユーザー更新エラー:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update user"})
		}

		return c.NoContent(http.StatusOK)
	}
}
