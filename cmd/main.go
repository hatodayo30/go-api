package main

import (
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"go-college/internal/infrastructure/db"
	"go-college/internal/infrastructure/repository"
	"go-college/internal/interface/handler"
	"go-college/internal/router"
	"go-college/internal/usecase"
)

func main() {
	// コマンドラインフラグの解析
	addr := flag.String("addr", ":8080", "tcp host:port to connect")
	flag.Parse()

	e := echo.New()

	// ミドルウェア設定
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: []string{"Content-Type,Accept,Origin,x-token"},
	}))

	// データベース初期化
	if err := db.InitDB(); err != nil {
		log.Fatalf("DB初期化に失敗しました: %v", err)
	}
	database := db.GetDB()

	// リポジトリの初期化（依存関係の注入）
	userRepo := repository.NewUserRepository(database)
	collectionRepo := repository.NewCollectionRepository(database)
	rankingRepo := repository.NewRankingRepository(database)
	finishRepo := repository.NewFinishRepository(database)
	gachaRepo := repository.NewGachaRepository(database) // ✅ GachaRepository 追加

	// ユースケースの初期化
	userUsecase := usecase.NewUserUsecase(userRepo)
	gameUsecase := usecase.NewGameUsecase(finishRepo)
	collectionUsecase := usecase.NewCollectionUsecase(collectionRepo)
	gachaUsecase := usecase.NewGachaUsecase(gachaRepo) // ✅ GachaUsecase 追加
	rankingUsecase := usecase.NewRankingUsecase(rankingRepo)

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(userUsecase)
	gameHandler := handler.NewGameHandler(gameUsecase)
	collectionHandler := handler.NewCollectionHandler(collectionUsecase)
	rankingHandler := handler.NewRankingHandler(rankingUsecase)
	gachaHandler := handler.NewGachaHandler(gachaUsecase) // ✅ GachaHandler 追加

	// ルーティング設定
	router.SetupRoutes(e, gachaHandler, userHandler, gameHandler, collectionHandler, rankingHandler, userRepo) // ✅ gachaHandler を追加

	// サーバー起動
	log.Printf("Server running on %s...", *addr)
	if err := e.Start(*addr); err != nil {
		log.Fatalf("failed to start server: %+v", err)
	}
}
