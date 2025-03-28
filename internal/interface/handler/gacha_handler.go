package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"go-college/internal/context/auth"
	"go-college/internal/interface/dto"
	"go-college/internal/usecase" // ✅ ユースケースを参照
)

type GachaHandler struct {
	gachaUsecase usecase.GachaUsecase
}

func NewGachaHandler(gachaUsecase usecase.GachaUsecase) *GachaHandler {
	return &GachaHandler{gachaUsecase: gachaUsecase}
}

// HandleGachaDraw ガチャを引くAPI
func (h *GachaHandler) HandleGachaDraw() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.GachaDrawRequest)
		if err := c.Bind(req); err != nil {
			log.Println("❌ [ERROR] バインドに失敗:", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストのバインドに失敗しました"})
		}

		// ガチャ回数のバリデーション
		if req.Times < 1 || req.Times > 10 {
			log.Println("❌ [ERROR] ガチャ回数エラー: ", req.Times)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "ガチャ回数は1以上10以下にしてください"})
		}

		ctx := c.Request().Context()
		userID, ok := auth.GetUserIDFromContext(ctx)
		if !ok || userID == "" {
			log.Println("❌ [ERROR] userID が空です")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "ユーザー認証が必要です"})
		}

		// **GachaUsecase を通じてガチャを実行**
		items, err := h.gachaUsecase.ExecuteGacha(userID, req.Times)
		if err != nil {
			log.Println("❌ [ERROR] ガチャ処理失敗:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ガチャの処理に失敗しました"})
		}

		res := dto.GachaDrawResponse{
			Results: items,
		}

		log.Println("✅ [INFO] ガチャ成功:", items)
		return c.JSON(http.StatusOK, res)
	}
}
