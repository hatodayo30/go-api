package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Conn 各リポジトリで利用するDB接続情報
var Conn *sql.DB

// DBConfig データベース設定
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// NewDB データベース接続を作成
func NewDB(config DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// SetDB グローバルなDBインスタンスを設定
func SetDB(db *sql.DB) {
	Conn = db
}

// LoadDBConfig 環境変数から `DBConfig` をロード
func LoadDBConfig() DBConfig {
	return DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}
}

func InitDB() error {
	config := LoadDBConfig()
	db, err := NewDB(config)
	if err != nil {
		return err
	}
	SetDB(db)
	return nil
}

// GetDB グローバルな DB 接続を取得
func GetDB() *sql.DB {
	return Conn
}
