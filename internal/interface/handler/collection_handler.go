package handler

import (
	"go-college/internal/context/auth"
	"go-college/internal/interface/dto"
	"go-college/internal/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CollectionHandler struct {
	collectionUsecase usecase.CollectionUsecase
}

func NewCollectionHandler(collectionUsecase usecase.CollectionUsecase) *CollectionHandler {
	return &CollectionHandler{collectionUsecase: collectionUsecase}
}

// HandleCollectionList ユーザーのコレクションアイテム一覧を取得
func (h *CollectionHandler) HandleCollectionList() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID, ok := auth.GetUserIDFromContext(ctx)
		log.Println("取得した userID", userID)
		if !ok || userID == "" {
			log.Println("❌ [ERROR] ユーザー認証が必要です")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "ユーザー認証が必要です"})
		}

		items, ownedCollections, totalCollections, err := h.collectionUsecase.GetUserCollectionList(userID)
		if err != nil {
			log.Printf("❌ [ERROR] ユーザーコレクション取得エラー: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "アイテムを取得できませんでした"})
		}
		log.Println("✅ [INFO] ユーザーコレクション取得成功")

		// レスポンス用の構造体に変換
		response := dto.CollectionResponse{
			Collections:      items,
			OwnedCollections: ownedCollections,
			TotalCollections: totalCollections,
		}

		return c.JSON(http.StatusOK, response)
	}
}
