package router

import (
	"go-college/internal/domain/repository"
	"go-college/internal/infrastructure/middleware"
	"go-college/internal/interface/handler"

	"github.com/labstack/echo/v4"
)

// SetupRoutes はアプリケーションのルーティングを設定する
func SetupRoutes(e *echo.Echo, gachaHandler *handler.GachaHandler, userHandler *handler.UserHandler, gameHandler *handler.GameHandler, collectionHandler *handler.CollectionHandler, rankingHandler *handler.RankingHandler, userRepo repository.UserRepository) {
	// 認証不要API
	e.GET("/setting/get", userHandler.HandleSettingGet())
	e.POST("/user/create", userHandler.HandleUserCreate())

	// 認証が必要なAPI
	authAPI := e.Group("", middleware.AuthenticateMiddleware(userRepo))
	authAPI.GET("/user/get", userHandler.HandleUserGet())
	authAPI.POST("/user/update", userHandler.HandleUserUpdate())
	authAPI.GET("/collection/list", collectionHandler.HandleCollectionList())
	authAPI.GET("/ranking/list", rankingHandler.HandleRankingList())
	authAPI.POST("/game/finish", gameHandler.HandleGameFinish())
	authAPI.POST("/gacha/draw", gachaHandler.HandleGachaDraw())
}
